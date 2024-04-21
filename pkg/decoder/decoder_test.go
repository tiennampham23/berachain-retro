package decoder

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
	"github.com/tiennampham23/berachain-airdrop/pkg/dapp/honey"
	"testing"
)

func Test_Decode(t *testing.T) {

	abi := honey.HoneyGeneratedMetaData.ABI

	data, err := hexutil.Decode("0xc6c3bbe60000000000000000000000003122c65da0e288fb745f07d8c81b10427b28e7ad0000000000000000000000006581e59a1c8da66ed0d313a0d4029dce2f746cc50000000000000000000000000000000000000000000000009f93707566a7c5cf")
	require.NoError(t, err)
	p, err := Decode(abi, data)
	require.NoError(t, err)

	t.Log(p)
}
