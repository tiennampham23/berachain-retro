package eth

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"math/big"
)

func SendTx(c *ethclient.Client, privateKey *ecdsa.PrivateKey, data []byte, value *big.Int, to common.Address) (*types.Transaction, error) {
	var (
		fromAddress, _ = GetAddressFromPrivateKey(privateKey)
	)
	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Logger().Errorw("Failed to get pending nonce", "err", err)
		return nil, err
	}

	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		log.Logger().Errorw("Failed to get gas price", "err", err)
		return nil, err
	}

	chainId, err := c.ChainID(context.Background())
	if err != nil {
		log.Logger().Errorw("Failed to get chain id", "err", err)
		return nil, err
	}

	tip := big.NewInt(1e9)

	gas, err := c.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  fromAddress,
		To:    &to,
		Value: value,
		Data:  data,
	})
	if err != nil {
		log.Logger().Errorw("Failed to estimate gas", "err", err)
		return nil, err
	}

	// Sign the transaction
	signedTx, err := types.SignNewTx(privateKey, types.LatestSignerForChainID(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasFeeCap: gasPrice,
		GasTipCap: tip,
		Gas:       gas,
		To:        &to,
		Value:     value,
		Data:      data,
	})
	if err != nil {
		log.Logger().Errorw("Failed to sign transaction", "err", err)
		return nil, err
	}

	// Send the transaction
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Logger().Errorw("Failed to send transaction", "err", err)
		return nil, err
	}

	return signedTx, nil
}
