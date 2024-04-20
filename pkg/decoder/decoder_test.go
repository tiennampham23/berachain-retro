package decoder

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
	"github.com/tiennampham23/berachain-airdrop/pkg/dapp/bex"
	"testing"
)

func Test_Decode(t *testing.T) {

	abi := bex.BexGeneratedMetaData.ABI

	data, err := hexutil.Decode("0xc02e8929000000000000000000000000a88572f08f79d28b8f864350f122c1cc0abb0d96")
	require.NoError(t, err)
	p, err := Decode(abi, data)
	require.NoError(t, err)

	t.Log(p)
}
