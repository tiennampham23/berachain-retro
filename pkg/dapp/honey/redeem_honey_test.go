package honey

import (
	"github.com/stretchr/testify/require"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
	"testing"
)

func Test_RedeemHoney(t *testing.T) {
	from, err := eth.GetAddressFromPrivateKey(privateKey)
	require.NoError(t, err)

	t.Log("Address", "from", from)

	amount := new(big.Int).SetUint64(3399088027719495625)
	tx, err := RedeemHoney(c, privateKey, STGUSDC, amount)
	require.NoError(t, err)

	t.Log("Submitted redeem honey transaction:", "hash", tx.Hash().Hex())
}
