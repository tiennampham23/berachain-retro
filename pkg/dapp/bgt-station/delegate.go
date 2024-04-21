package bgt_station

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
	ParsedStakingABI, _ = StakingGeneratedModuleMetaData.GetAbi()
)

var (
	StakingContractAddress = common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF")
)

var (
	Validators = []common.Address{
		common.HexToAddress("0x032238ba76Aadaa7C891967c4491fC18f81C6189"),
	}
)

func Delegate(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	validator common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	//
	// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
	//
	// Solidity: function delegate(address validatorAddress, uint256 amount) payable returns(bool)
	// func (_StakingGeneratedModule *StakingGeneratedModuleTransactor) Delegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	//	return _StakingGeneratedModule.contract.Transact(opts, "delegate", validatorAddress, amount)
	// }

	payload, err := ParsedStakingABI.Pack("delegate", validator, amount)
	if err != nil {
		return nil, err
	}

	log.Logger().Infow("Swap info", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, nil, StakingContractAddress)
	if err != nil {
		log.Logger().Errorw("Failed to send delegate tx", "error", err)
		return nil, err
	}

	return tx, nil
}
