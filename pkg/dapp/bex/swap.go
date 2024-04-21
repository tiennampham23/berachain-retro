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
	dappCommon "github.com/tiennampham23/berachain-airdrop/pkg/dapp/common"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
	"time"
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
		err := dappCommon.ApproveIfNeed(c, privateKey, tokenIn, amount, Erc20ModuleContract)
		if err != nil {
			log.Logger().Errorw("Approve failed", "err", err)
			return nil, err
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

	batchSwapPayload, err := bexABI.Pack("batchSwap", batchSwapKind, batchStepSwaps, deadline)
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
