package bend

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"os"
	"testing"
)

var (
	c, _          = ethclient.Dial(os.Getenv("RPC_URL"))
	privateKey, _ = crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
)

func Test_Supply(t *testing.T) {
	tx, err := Supply(c, privateKey, common.HexToAddress("0x8239FBb3e3D0C2cDFd7888D8aF7701240Ac4DcA4"), new(big.Int).SetInt64(457588966126130))
	require.NoError(t, err)

	t.Log("Supply tx:", tx.Hash().Hex())
}

func Test_Borrow(t *testing.T) {
	tx, err := Borrow(c, privateKey, common.HexToAddress("0x7EeCA4205fF31f947EdBd49195a7A88E6A91161B"), new(big.Int).SetInt64(616143412629750908))
	require.NoError(t, err)

	t.Log("Borrow tx:", tx.Hash().Hex())
}

func Test_Withdraw(t *testing.T) {
	tx, err := Withdraw(c, privateKey, common.HexToAddress("0x7EeCA4205fF31f947EdBd49195a7A88E6A91161B"), new(big.Int).SetInt64(157073566539289850))
	require.NoError(t, err)

	t.Log("Withdraw tx:", tx.Hash().Hex())
}

func Test_Repay(t *testing.T) {
	tx, err := Repay(c, privateKey, common.HexToAddress("0x7EeCA4205fF31f947EdBd49195a7A88E6A91161B"), new(big.Int).SetInt64(316143453324675217))
	require.NoError(t, err)

	t.Log("Withdraw tx:", tx.Hash().Hex())
}
