package bex

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"os"
	"testing"
)

var (
	c, _          = ethclient.Dial(os.Getenv("RPC_URL"))
	privateKey, _ = crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
)

func Test_SwapNative(t *testing.T) {
	tokenIn := Bear
	tokenOut := WBear
	amount := new(big.Int).SetInt64(30000000000000000)

	tx, err := Swap(c, privateKey, tokenIn, tokenOut, amount)
	require.NoError(t, err)

	t.Log("Swap tx:", tx.Hash().Hex())

}

func Test_SwapERC20(t *testing.T) {
	tokenIn := Honey
	tokenOut := WBear
	amount := new(big.Int).SetInt64(7371211818077651176)

	tx, err := Swap(c, privateKey, tokenIn, tokenOut, amount)
	require.NoError(t, err)

	t.Log("Swap tx:", tx.Hash().Hex())

}
