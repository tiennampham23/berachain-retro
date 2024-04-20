package bex

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func Test_AddLiquidityForSingleToken(t *testing.T) {
	tokenIn := WBear
	amount := new(big.Int).SetInt64(10000000000000000)

	poolAddress := common.HexToAddress("0xa88572F08f79D28b8f864350f122c1CC0AbB0d96")

	tx, err := AddLiquidityForSingleToken(c, privateKey, tokenIn, poolAddress, amount)
	require.NoError(t, err)

	t.Log("AddLiquidity tx:", tx.Hash().Hex())

}
