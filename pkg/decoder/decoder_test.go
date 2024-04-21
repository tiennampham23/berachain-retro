package decoder

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
	"github.com/tiennampham23/berachain-airdrop/pkg/dapp/bend"
	"testing"
)

func Test_Decode(t *testing.T) {

	abi := bend.BendGeneratedMetaData.ABI

	data, err := hexutil.Decode("0x573ade810000000000000000000000007eeca4205ff31f947edbd49195a7a88e6a91161b00000000000000000000000000000000000000000000000004632acbb97b889100000000000000000000000000000000000000000000000000000000000000020000000000000000000000004c92939202935e1327a387339d62c4b3f5e7bf6a")
	require.NoError(t, err)
	p, err := Decode(abi, data)
	require.NoError(t, err)

	t.Log(p)
}
