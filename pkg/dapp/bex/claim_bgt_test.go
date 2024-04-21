package bex

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ClaimBGT(t *testing.T) {
	poolAddress := common.HexToAddress("0xa88572F08f79D28b8f864350f122c1CC0AbB0d96")

	tx, err := WithdrawAllDepositorRewards(
		c,
		privateKey,
		poolAddress,
	)
	require.NoError(t, err)

	if tx == nil {
		t.Log("Tx is nil")
		return
	}
	t.Log("tx:", tx.Hash().Hex())
}
