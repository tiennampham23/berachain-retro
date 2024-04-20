package bex

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/erc20"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
	"time"
)

var (
	client *resty.Client
)

var (
	WBear                     = common.HexToAddress("0x5806E416dA447b267cEA759358cF22Cc41FAE80F")
	Honey                     = common.HexToAddress("0x7EeCA4205fF31f947EdBd49195a7A88E6A91161B")
	Bear                      = common.HexToAddress("0x0000000000000000000000000000000000000000")
	Erc20ModuleContract       = common.HexToAddress("0x0000000000000000000000000000000000696969")
	parsedABI, _              = BexGeneratedMetaData.GetAbi()
	batchSwapKind       uint8 = 0
	ApproveMax, _             = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
	SwapContractAddress       = common.HexToAddress("0x0d5862FDbdd12490f9b4De54c236cff63B038074")
)

type Step struct {
	Pool      string `json:"pool"`
	AssetIn   string `json:"assetIn"`
	AmountIn  string `json:"amountIn"`
	AssetOut  string `json:"assetOut"`
	AmountOut string `json:"amountOut"`
}

type Response struct {
	Steps []Step `json:"steps"`
}

func init() {
	client = resty.New().SetTimeout(10 * time.Second)
}

func getDexRoute(baseAsset string, quoteAsset string, amount string) (Response, error) {
	var response Response
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"quoteAsset": quoteAsset,
			"baseAsset":  baseAsset,
			"amount":     amount,
			"swap_type":  "given_in",
		}).
		SetResult(&response).
		Get("https://artio-80085-dex-router.berachain.com/dex/route")

	if err != nil {
		return Response{}, err
	}

	if resp.IsError() {
		return Response{}, fmt.Errorf("%s", resp.String())
	}

	return response, nil
}

func Swap(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	tokenIn common.Address,
	tokenOut common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	var (
		baseToken  = tokenIn
		quoteToken = tokenOut
		from, _    = eth.GetAddressFromPrivateKey(privateKey)
		value      = new(big.Int)
	)

	if tokenIn.Cmp(Bear) == 0 {
		baseToken = WBear
		value = amount
	}

	if tokenIn.Cmp(Bear) == 0 && tokenOut.Cmp(WBear) == 0 {
		wrapPayload, err := Wrap()
		if err != nil {
			return nil, err
		}

		tx, err := eth.SendTx(c, privateKey, wrapPayload, value, WBear)
		if err != nil {
			log.Logger().Errorw("Wrap BEAR failed", "err", err)
			return nil, err
		}
		log.Logger().Infow("Wrap BEAR successfully", "tx", tx.Hash().Hex())
		return tx, nil
	}

	routeResp, err := getDexRoute(
		baseToken.String(),
		quoteToken.String(),
		amount.String(),
	)

	if err != nil {
		log.Logger().Errorw("Failed to get Dex route", "error", err)
		return nil, err
	}

	if len(routeResp.Steps) == 0 {
		return nil, fmt.Errorf("no steps found")
	}

	if tokenIn.Cmp(Bear) == 0 {
		// check balance
		balance, err := c.BalanceAt(context.Background(), from, nil)
		if err != nil {
			return nil, err
		}
		if balance.Cmp(amount) < 0 {
			return nil, errors.New("insufficient balance")
		}
	} else {
		balance, err := erc20.BalanceOf(c, from, tokenIn)
		if err != nil {
			log.Logger().Errorw("Failed to get balance", "error", err)
			return nil, err
		}

		if balance.Cmp(amount) < 0 {
			return nil, errors.New("insufficient balance")
		}
		allowance, err := erc20.Allowance(c, from, Erc20ModuleContract, tokenIn)
		if err != nil {
			return nil, err
		}

		if allowance.Cmp(amount) < 0 {
			approvePayload, err := erc20.Approve(Erc20ModuleContract, ApproveMax)
			if err != nil {
				log.Logger().Errorw("Failed to approve payload", "error", err)
				return nil, err
			}

			tx, err := eth.SendTx(c, privateKey, approvePayload, nil, tokenIn)
			if err != nil {
				log.Logger().Errorw("Failed to send approval tx", "error", err)
				return nil, err
			}
			log.Logger().Infow("Approve tx", "tx", tx.Hash())

			for {
				_, isPending, err := c.TransactionByHash(context.Background(), tx.Hash())
				if err != nil {
					log.Logger().Errorw("Failed to get approve transaction by hash", "error", err)
					time.Sleep(5 * time.Second)
					continue
				}
				if isPending {
					time.Sleep(5 * time.Second)
					continue
				}
				break
			}
		}
	}

	deadline := new(big.Int).SetInt64(time.Now().Add(time.Minute).Unix())

	batchStepSwaps := make([]IERC20DexModuleBatchSwapStep, 0)

	for index, step := range routeResp.Steps {
		amountIn, _ := new(big.Int).SetString(step.AmountIn, 10)

		amountOut := big.NewInt(0)

		// only adjust amountOut on the last step
		if index+1 == len(routeResp.Steps) {
			lastAmountOut, _ := new(big.Int).SetString(step.AmountOut, 10)
			amountOut = new(big.Int).Div(lastAmountOut, new(big.Int).SetInt64(2))
			log.Logger().Infow("Amount out", "out", amountOut.String(), "lastAmountOut", step.AmountOut)

		}

		batchStepSwaps = append(batchStepSwaps, IERC20DexModuleBatchSwapStep{
			PoolId:    common.HexToAddress(step.Pool),
			AssetIn:   tokenIn,
			AmountIn:  amountIn,
			AssetOut:  tokenOut,
			AmountOut: amountOut,
			UserData:  []byte{},
		})
	}

	batchSwapPayload, err := parsedABI.Pack("batchSwap", batchSwapKind, batchStepSwaps, deadline)
	if err != nil {
		log.Logger().Fatalf("Failed to pack transaction data: %v", err)
	}

	log.Logger().Infow("Swap info", "payload", hexutil.Encode(batchSwapPayload))

	swapTx, err := eth.SendTx(c, privateKey, batchSwapPayload, value, SwapContractAddress)
	if err != nil {
		log.Logger().Errorw("Failed to send swap tx", "error", err)
		return nil, err
	}

	return swapTx, nil
}
