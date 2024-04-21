package beranames

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
)

var (
	BeranameContractAddress = common.HexToAddress("0x8D20B92B4163140F413AA52A4106fF9490bf2122")
	parsedABI, _            = BeranameGeneratedMetaData.GetAbi()
)

func MintBeraname(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	chars []string,
) (*types.Transaction, error) {
	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)

	// MintNative is a paid mutator transaction binding the contract method 0x3d30c7f6.
	//
	// Solidity: function mintNative(string[] chars, uint256 duration, address whois, string metadataURI, address to) payable returns(uint256)
	// func (_BeranameGenerated *BeranameGeneratedTransactor) MintNative(opts *bind.TransactOpts, chars []string, duration *big.Int, whois common.Address, metadataURI string, to common.Address) (*types.Transaction, error) {
	//	 return _BeranameGenerated.contract.Transact(opts, "mintNative", chars, duration, whois, metadataURI, to)
	// }

	payload, err := parsedABI.Pack("mintNative", chars, new(big.Int).SetInt64(1), from, "https://beranames.com/api/metadata/69", from)
	if err != nil {
		log.Logger().Errorw("failed to pack payload", "err", err)
		return nil, err
	}

	log.Logger().Infow("Mint payload", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, new(big.Int).SetInt64(450158455776433), BeranameContractAddress)
	if err != nil {
		log.Logger().Errorw("failed to send mint", "error", err)
		return nil, err
	}

	return tx, nil
}
