package honey

import (
	"github.com/stretchr/testify/require"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
	"testing"
)

func Test_MintHoney(t *testing.T) {
	from, err := eth.GetAddressFromPrivateKey(privateKey)
	require.NoError(t, err)

	t.Log("Address", "from", from)

	amount := new(big.Int).SetUint64(3416168872079895101)
	tx, err := Mint(c, privateKey, STGUSDC, amount)
	require.NoError(t, err)

	t.Log("Submitted transaction:", "hash", tx.Hash().Hex())
}
