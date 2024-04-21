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

func Undelegate(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	validator common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
	//
	// Solidity: function undelegate(address validatorAddress, uint256 amount) payable returns(bool)
	// func (_StakingGeneratedModule *StakingGeneratedModuleTransactor) Undelegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	//	 return _StakingGeneratedModule.contract.Transact(opts, "undelegate", validatorAddress, amount)
	// }

	payload, err := ParsedStakingABI.Pack("undelegate", validator, amount)
	if err != nil {
		return nil, err
	}

	log.Logger().Infow("Undelegate info", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, nil, StakingContractAddress)
	if err != nil {
		log.Logger().Errorw("Failed to send undelegate tx", "error", err)
		return nil, err
	}

	return tx, nil
}
