package bgt_station

import (
	"github.com/stretchr/testify/require"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
	"testing"
)

func Test_Delegate(t *testing.T) {
	from, err := eth.GetAddressFromPrivateKey(privateKey)
	require.NoError(t, err)

	t.Log("Address", "from", from)

	amount := new(big.Int).SetUint64(0)

	validator := Validators[0]
	tx, err := Delegate(c, privateKey, validator, amount)
	require.NoError(t, err)

	t.Log("Submitted transaction:", "hash", tx.Hash().Hex())
}
