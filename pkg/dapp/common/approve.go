package common

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/erc20"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
	"time"
)

var (
	ApproveMax, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
)

func ApproveIfNeed(c *ethclient.Client, privateKey *ecdsa.PrivateKey, tokenIn common.Address, amount *big.Int, contract common.Address) error {
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
	allowance, err := erc20.Allowance(c, from, contract, tokenIn)
	if err != nil {
		return err
	}

	if allowance.Cmp(amount) < 0 {
		approvePayload, err := erc20.Approve(contract, ApproveMax)
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
