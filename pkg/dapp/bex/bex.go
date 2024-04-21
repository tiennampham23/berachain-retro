package bex

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/erc20"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
	"time"
)

var (
	WBear                        = common.HexToAddress("0x5806E416dA447b267cEA759358cF22Cc41FAE80F")
	Honey                        = common.HexToAddress("0x7EeCA4205fF31f947EdBd49195a7A88E6A91161B")
	Bear                         = common.HexToAddress("0x0000000000000000000000000000000000000000")
	Erc20ModuleContract          = common.HexToAddress("0x0000000000000000000000000000000000696969")
	bexABI, _                    = BexGeneratedMetaData.GetAbi()
	batchSwapKind          uint8 = 0
	ApproveMax, _                = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
	SwapContractAddress          = common.HexToAddress("0x0d5862FDbdd12490f9b4De54c236cff63B038074")
	RewardModuleContract         = common.HexToAddress("0x55684e2cA2bace0aDc512C1AFF880b15b8eA7214")
	MinRewardAmountToClaim       = new(big.Int).SetInt64(0)
)

var (
	client *resty.Client
)

func init() {
	client = resty.New().SetTimeout(10 * time.Second)
}

func ApproveIfNeed(c *ethclient.Client, privateKey *ecdsa.PrivateKey, tokenIn common.Address, amount *big.Int) error {
	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)
	balance, err := erc20.BalanceOf(c, from, tokenIn)
	if err != nil {
		log.Logger().Errorw("Failed to get balance", "error", err)
		return err
	}

	if balance.Cmp(amount) < 0 {
		return errors.New("insufficient balance")
	}
	allowance, err := erc20.Allowance(c, from, Erc20ModuleContract, tokenIn)
	if err != nil {
		return err
	}

	if allowance.Cmp(amount) < 0 {
		approvePayload, err := erc20.Approve(Erc20ModuleContract, ApproveMax)
		if err != nil {
			log.Logger().Errorw("Failed to approve payload", "error", err)
			return err
		}

		tx, err := eth.SendTx(c, privateKey, approvePayload, nil, tokenIn)
		if err != nil {
			log.Logger().Errorw("Failed to send approval tx", "error", err)
			return err
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

	return nil
}
