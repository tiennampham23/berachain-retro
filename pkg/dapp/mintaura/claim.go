package mintaura

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"

	"math/big"
)

var (
	parsedABI, _ = MintAuraGeneratedMetaData.GetAbi()
)

func Claim(
	receiver common.Address,
	quantity *big.Int,
	currency common.Address,
	pricePerToken *big.Int,
	allowlistProof IDropAllowlistProof,
	data []byte,
) ([]byte, error) {
	// Pack the data to send a transaction
	inputData, err := parsedABI.Pack("claim", receiver, quantity, currency, pricePerToken, allowlistProof, data)
	if err != nil {
		log.Logger().Fatalf("Failed to pack transaction data: %v", err)
	}

	return inputData, nil
}
