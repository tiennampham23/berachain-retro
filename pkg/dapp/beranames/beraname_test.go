package beranames

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

func Test_Mint(t *testing.T) {
	tx, err := MintBeraname(c, privateKey, []string{
		"B",
		"T",
		"C",
		"@",
		"2",
		"0",
		"2",
		"5",
	})
	require.NoError(t, err)

	t.Log("Mint Beranames tx:", tx.Hash().Hex())
}
