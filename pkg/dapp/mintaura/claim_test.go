package mintaura

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"math/big"
	"testing"
)

func Test_Claim(t *testing.T) {
	receiver := common.HexToAddress("0x84abd05Cb1Bc35d22eBb12322B7cCf2aDfC513d2")
	quantity := big.NewInt(1)
	currency := common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	pricePerToken := big.NewInt(50000000000000000)

	quantityLimitPerWallet, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	byte32 := [32]byte{}
	allowlistProof := IDropAllowlistProof{
		Proof:                  [][32]byte{byte32},
		QuantityLimitPerWallet: quantityLimitPerWallet,
		PricePerToken:          pricePerToken,
		Currency:               currency,
	}

	var value []byte

	data, err := Claim(
		receiver,
		quantity,
		currency,
		pricePerToken,
		allowlistProof,
		value,
	)

	require.NoError(t, err)

	log.Logger().Infow(hexutil.Encode(data))

	c, err := ethclient.Dial("https://artio.rpc.berachain.com")
	require.NoError(t, err)

	contractAddress := common.HexToAddress("0x6Dc544b917904959e5c26b9202907ca16BDcdb4A")
	gas, err := c.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  receiver,
		To:    &contractAddress,
		Value: pricePerToken,
		Data:  data,
	})
	require.NoError(t, err)

	t.Log("Gas used:", gas)
}
