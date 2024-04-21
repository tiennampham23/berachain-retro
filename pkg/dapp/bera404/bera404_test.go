package bera404

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var (
	c, _          = ethclient.Dial(os.Getenv("RPC_URL"))
	privateKey, _ = crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
)

func Test_MintBera404(t *testing.T) {
	tx, err := MintBera404(c, privateKey)
	require.NoError(t, err)

	t.Log("Mint bera404 tx:", tx.Hash().Hex())
}
