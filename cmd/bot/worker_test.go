package bot

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_GetTx(t *testing.T) {
	c, err := ethclient.Dial("https://artio.rpc.berachain.com")
	require.NoError(t, err)

	tx, _, err := c.TransactionByHash(context.Background(), common.HexToHash("0x73f9502d06f0bb891ec3cbce45b6de3c43bb0a645e991506eb6652a1e5b50bba"))
	require.NoError(t, err)

	t.Log(tx.Gas())

}
