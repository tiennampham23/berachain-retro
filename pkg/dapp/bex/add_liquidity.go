package bex

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func AddLiquidity(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	tokenIn common.Address,
	tokenOut common.Address,
	amount *big.Int,
) (*types.Transaction, error) {

}
