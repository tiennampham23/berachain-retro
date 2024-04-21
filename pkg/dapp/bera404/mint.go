package bera404

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
	ContractAddress = common.HexToAddress("0x1F136a43101D12F98c9887D46D7cDbEFACdd573D")
	parsedABI, _    = Bera404MetaData.GetAbi()
)

func MintBera404(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
) (*types.Transaction, error) {

	// BoobaMint is a paid mutator transaction binding the contract method 0x1f78fa9f.
	//
	// Solidity: function boobaMint(uint256 value_, bool mintNFT) returns()
	// func (_Bera404 *Bera404Transactor) BoobaMint(opts *bind.TransactOpts, value_ *big.Int, mintNFT bool) (*types.Transaction, error) {
	//	 return _Bera404.contract.Transact(opts, "boobaMint", value_, mintNFT)
	// }
	//
	payload, err := parsedABI.Pack("boobaMint", new(big.Int).SetUint64(1000000000000000000), false)
	if err != nil {
		log.Logger().Errorw("failed to pack payload", "err", err)
		return nil, err
	}

	log.Logger().Infow("Mint payload", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, nil, ContractAddress)
	if err != nil {
		log.Logger().Errorw("failed to send mint", "error", err)
		return nil, err
	}

	return tx, nil
}
