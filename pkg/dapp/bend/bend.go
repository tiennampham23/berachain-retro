package bend

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	dappCommon "github.com/tiennampham23/berachain-airdrop/pkg/dapp/common"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
)

var (
	parsedBendABI, _    = BendGeneratedMetaData.GetAbi()
	BendContractAddress = common.HexToAddress("0x9261b5891d3556e829579964B38fe706D0A2D04a")
	interestRateMode    = new(big.Int).SetInt64(2)
)

func Supply(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	tokenIn common.Address,
	amountIn *big.Int,
) (*types.Transaction, error) {
	err := dappCommon.ApproveIfNeed(c, privateKey, tokenIn, amountIn, BendContractAddress)
	if err != nil {
		return nil, err
	}

	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)

	// Supply is a paid mutator transaction binding the contract method 0x617ba037.
	//
	// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
	// func (_BendGenerated *BendGeneratedTransactor) Supply(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	//	 return _BendGenerated.contract.Transact(opts, "supply", asset, amount, onBehalfOf, referralCode)
	// }
	payload, err := parsedBendABI.Pack("supply", tokenIn, amountIn, from, uint16(0))
	if err != nil {
		log.Logger().Errorw("Failed to pack supply", "err", err)
		return nil, err
	}

	log.Logger().Infow("Supply payload", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, nil, BendContractAddress)
	if err != nil {
		log.Logger().Errorw("Failed to send supply tx", "err", err)
		return nil, err
	}

	return tx, nil
}

func Borrow(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	tokenIn common.Address,
	amountIn *big.Int,
) (*types.Transaction, error) {
	err := dappCommon.ApproveIfNeed(c, privateKey, tokenIn, amountIn, BendContractAddress)
	if err != nil {
		return nil, err
	}

	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)

	// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
	//
	// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
	// func (_BendGenerated *BendGeneratedTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	//	 return _BendGenerated.contract.Transact(opts, "borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
	// }
	payload, err := parsedBendABI.Pack("borrow", tokenIn, amountIn, interestRateMode, uint16(0), from)
	if err != nil {
		log.Logger().Errorw("Failed to pack borrow", "err", err)
		return nil, err
	}

	log.Logger().Infow("Borrow payload", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, nil, BendContractAddress)
	if err != nil {
		log.Logger().Errorw("Failed to send supply tx", "err", err)
		return nil, err
	}

	return tx, nil
}

func Withdraw(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	tokenOut common.Address,
	amount *big.Int,
) (*types.Transaction, error) {

	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)

	// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
	//
	// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
	//func (_BendGenerated *BendGeneratedTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	//	return _BendGenerated.contract.Transact(opts, "withdraw", asset, amount, to)
	//}
	payload, err := parsedBendABI.Pack("withdraw", tokenOut, amount, from)
	if err != nil {
		log.Logger().Errorw("Failed to pack withdraw", "err", err)
		return nil, err
	}

	log.Logger().Infow("Withdraw payload", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, nil, BendContractAddress)
	if err != nil {
		log.Logger().Errorw("Failed to send withdrawn tx", "err", err)
		return nil, err
	}

	return tx, nil
}

func Repay(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	tokenOut common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)

	// Repay is a paid mutator transaction binding the contract method 0x573ade81.
	//
	// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
	//func (_BendGenerated *BendGeneratedTransactor) Repay(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	//	return _BendGenerated.contract.Transact(opts, "repay", asset, amount, interestRateMode, onBehalfOf)
	//}
	payload, err := parsedBendABI.Pack("repay", tokenOut, amount, interestRateMode, from)
	if err != nil {
		log.Logger().Errorw("Failed to pack withdraw", "err", err)
		return nil, err
	}

	log.Logger().Infow("Repay payload", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, nil, BendContractAddress)
	if err != nil {
		log.Logger().Errorw("Failed to send repay tx", "err", err)
		return nil, err
	}

	return tx, nil
}
