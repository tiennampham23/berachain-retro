package erc20

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var (
	parsedABI, _ = Erc20GeneratedMetaData.GetAbi()
)

func Approve(contract common.Address, amount *big.Int) ([]byte, error) {
	payload, err := parsedABI.Pack("approve", contract, amount)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func Allowance(c *ethclient.Client, from common.Address, contract common.Address, tokenIn common.Address) (*big.Int, error) {
	// check allowance
	tokenContract, err := NewErc20Generated(tokenIn, c)
	if err != nil {
		return nil, err
	}

	allowance, err := tokenContract.Allowance(&bind.CallOpts{}, from, contract)
	if err != nil {
		return nil, err
	}

	return allowance, nil
}

func BalanceOf(c *ethclient.Client, from common.Address, tokenIn common.Address) (*big.Int, error) {
	// check allowance
	tokenContract, err := NewErc20Generated(tokenIn, c)
	if err != nil {
		return nil, err
	}

	balance, err := tokenContract.BalanceOf(&bind.CallOpts{}, from)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
