// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package honey

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ERC20HoneyERC20Exchangable is an auto generated low-level Go binding around an user-defined struct.
type ERC20HoneyERC20Exchangable struct {
	Collateral     common.Address
	Enabled        bool
	MintRate       *big.Int
	RedemptionRate *big.Int
}

// HoneyGeneratedMetaData contains all meta data concerning the HoneyGenerated contract.
var HoneyGeneratedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_honey\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"erc20Module\",\"outputs\":[{\"internalType\":\"contractIERC20Module\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangable\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"mintRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionRate\",\"type\":\"uint256\"}],\"internalType\":\"structERC20Honey.ERC20Exchangable[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"honey\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"honeyModule\",\"outputs\":[{\"internalType\":\"contractIHoneyModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HoneyGeneratedABI is the input ABI used to generate the binding from.
// Deprecated: Use HoneyGeneratedMetaData.ABI instead.
var HoneyGeneratedABI = HoneyGeneratedMetaData.ABI

// HoneyGenerated is an auto generated Go binding around an Ethereum contract.
type HoneyGenerated struct {
	HoneyGeneratedCaller     // Read-only binding to the contract
	HoneyGeneratedTransactor // Write-only binding to the contract
	HoneyGeneratedFilterer   // Log filterer for contract events
}

// HoneyGeneratedCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoneyGeneratedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoneyGeneratedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoneyGeneratedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoneyGeneratedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoneyGeneratedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoneyGeneratedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoneyGeneratedSession struct {
	Contract     *HoneyGenerated   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoneyGeneratedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoneyGeneratedCallerSession struct {
	Contract *HoneyGeneratedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HoneyGeneratedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoneyGeneratedTransactorSession struct {
	Contract     *HoneyGeneratedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HoneyGeneratedRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoneyGeneratedRaw struct {
	Contract *HoneyGenerated // Generic contract binding to access the raw methods on
}

// HoneyGeneratedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoneyGeneratedCallerRaw struct {
	Contract *HoneyGeneratedCaller // Generic read-only contract binding to access the raw methods on
}

// HoneyGeneratedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoneyGeneratedTransactorRaw struct {
	Contract *HoneyGeneratedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoneyGenerated creates a new instance of HoneyGenerated, bound to a specific deployed contract.
func NewHoneyGenerated(address common.Address, backend bind.ContractBackend) (*HoneyGenerated, error) {
	contract, err := bindHoneyGenerated(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoneyGenerated{HoneyGeneratedCaller: HoneyGeneratedCaller{contract: contract}, HoneyGeneratedTransactor: HoneyGeneratedTransactor{contract: contract}, HoneyGeneratedFilterer: HoneyGeneratedFilterer{contract: contract}}, nil
}

// NewHoneyGeneratedCaller creates a new read-only instance of HoneyGenerated, bound to a specific deployed contract.
func NewHoneyGeneratedCaller(address common.Address, caller bind.ContractCaller) (*HoneyGeneratedCaller, error) {
	contract, err := bindHoneyGenerated(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoneyGeneratedCaller{contract: contract}, nil
}

// NewHoneyGeneratedTransactor creates a new write-only instance of HoneyGenerated, bound to a specific deployed contract.
func NewHoneyGeneratedTransactor(address common.Address, transactor bind.ContractTransactor) (*HoneyGeneratedTransactor, error) {
	contract, err := bindHoneyGenerated(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoneyGeneratedTransactor{contract: contract}, nil
}

// NewHoneyGeneratedFilterer creates a new log filterer instance of HoneyGenerated, bound to a specific deployed contract.
func NewHoneyGeneratedFilterer(address common.Address, filterer bind.ContractFilterer) (*HoneyGeneratedFilterer, error) {
	contract, err := bindHoneyGenerated(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoneyGeneratedFilterer{contract: contract}, nil
}

// bindHoneyGenerated binds a generic wrapper to an already deployed contract.
func bindHoneyGenerated(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HoneyGeneratedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoneyGenerated *HoneyGeneratedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HoneyGenerated.Contract.HoneyGeneratedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoneyGenerated *HoneyGeneratedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoneyGenerated.Contract.HoneyGeneratedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoneyGenerated *HoneyGeneratedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoneyGenerated.Contract.HoneyGeneratedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoneyGenerated *HoneyGeneratedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HoneyGenerated.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoneyGenerated *HoneyGeneratedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoneyGenerated.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoneyGenerated *HoneyGeneratedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoneyGenerated.Contract.contract.Transact(opts, method, params...)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_HoneyGenerated *HoneyGeneratedCaller) Erc20Module(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HoneyGenerated.contract.Call(opts, &out, "erc20Module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_HoneyGenerated *HoneyGeneratedSession) Erc20Module() (common.Address, error) {
	return _HoneyGenerated.Contract.Erc20Module(&_HoneyGenerated.CallOpts)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_HoneyGenerated *HoneyGeneratedCallerSession) Erc20Module() (common.Address, error) {
	return _HoneyGenerated.Contract.Erc20Module(&_HoneyGenerated.CallOpts)
}

// Honey is a free data retrieval call binding the contract method 0x36b2c4b2.
//
// Solidity: function honey() view returns(address)
func (_HoneyGenerated *HoneyGeneratedCaller) Honey(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HoneyGenerated.contract.Call(opts, &out, "honey")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Honey is a free data retrieval call binding the contract method 0x36b2c4b2.
//
// Solidity: function honey() view returns(address)
func (_HoneyGenerated *HoneyGeneratedSession) Honey() (common.Address, error) {
	return _HoneyGenerated.Contract.Honey(&_HoneyGenerated.CallOpts)
}

// Honey is a free data retrieval call binding the contract method 0x36b2c4b2.
//
// Solidity: function honey() view returns(address)
func (_HoneyGenerated *HoneyGeneratedCallerSession) Honey() (common.Address, error) {
	return _HoneyGenerated.Contract.Honey(&_HoneyGenerated.CallOpts)
}

// HoneyModule is a free data retrieval call binding the contract method 0xbabc04fc.
//
// Solidity: function honeyModule() view returns(address)
func (_HoneyGenerated *HoneyGeneratedCaller) HoneyModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HoneyGenerated.contract.Call(opts, &out, "honeyModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HoneyModule is a free data retrieval call binding the contract method 0xbabc04fc.
//
// Solidity: function honeyModule() view returns(address)
func (_HoneyGenerated *HoneyGeneratedSession) HoneyModule() (common.Address, error) {
	return _HoneyGenerated.Contract.HoneyModule(&_HoneyGenerated.CallOpts)
}

// HoneyModule is a free data retrieval call binding the contract method 0xbabc04fc.
//
// Solidity: function honeyModule() view returns(address)
func (_HoneyGenerated *HoneyGeneratedCallerSession) HoneyModule() (common.Address, error) {
	return _HoneyGenerated.Contract.HoneyModule(&_HoneyGenerated.CallOpts)
}

// PreviewMint is a free data retrieval call binding the contract method 0xd1f810a5.
//
// Solidity: function previewMint(address collateral, uint256 amount) view returns(uint256)
func (_HoneyGenerated *HoneyGeneratedCaller) PreviewMint(opts *bind.CallOpts, collateral common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HoneyGenerated.contract.Call(opts, &out, "previewMint", collateral, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xd1f810a5.
//
// Solidity: function previewMint(address collateral, uint256 amount) view returns(uint256)
func (_HoneyGenerated *HoneyGeneratedSession) PreviewMint(collateral common.Address, amount *big.Int) (*big.Int, error) {
	return _HoneyGenerated.Contract.PreviewMint(&_HoneyGenerated.CallOpts, collateral, amount)
}

// PreviewMint is a free data retrieval call binding the contract method 0xd1f810a5.
//
// Solidity: function previewMint(address collateral, uint256 amount) view returns(uint256)
func (_HoneyGenerated *HoneyGeneratedCallerSession) PreviewMint(collateral common.Address, amount *big.Int) (*big.Int, error) {
	return _HoneyGenerated.Contract.PreviewMint(&_HoneyGenerated.CallOpts, collateral, amount)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0xcbe52ae3.
//
// Solidity: function previewRedeem(address collateral, uint256 amount) view returns(uint256)
func (_HoneyGenerated *HoneyGeneratedCaller) PreviewRedeem(opts *bind.CallOpts, collateral common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HoneyGenerated.contract.Call(opts, &out, "previewRedeem", collateral, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0xcbe52ae3.
//
// Solidity: function previewRedeem(address collateral, uint256 amount) view returns(uint256)
func (_HoneyGenerated *HoneyGeneratedSession) PreviewRedeem(collateral common.Address, amount *big.Int) (*big.Int, error) {
	return _HoneyGenerated.Contract.PreviewRedeem(&_HoneyGenerated.CallOpts, collateral, amount)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0xcbe52ae3.
//
// Solidity: function previewRedeem(address collateral, uint256 amount) view returns(uint256)
func (_HoneyGenerated *HoneyGeneratedCallerSession) PreviewRedeem(collateral common.Address, amount *big.Int) (*big.Int, error) {
	return _HoneyGenerated.Contract.PreviewRedeem(&_HoneyGenerated.CallOpts, collateral, amount)
}

// GetExchangable is a paid mutator transaction binding the contract method 0xd20c48a8.
//
// Solidity: function getExchangable() returns((address,bool,uint256,uint256)[])
func (_HoneyGenerated *HoneyGeneratedTransactor) GetExchangable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoneyGenerated.contract.Transact(opts, "getExchangable")
}

// GetExchangable is a paid mutator transaction binding the contract method 0xd20c48a8.
//
// Solidity: function getExchangable() returns((address,bool,uint256,uint256)[])
func (_HoneyGenerated *HoneyGeneratedSession) GetExchangable() (*types.Transaction, error) {
	return _HoneyGenerated.Contract.GetExchangable(&_HoneyGenerated.TransactOpts)
}

// GetExchangable is a paid mutator transaction binding the contract method 0xd20c48a8.
//
// Solidity: function getExchangable() returns((address,bool,uint256,uint256)[])
func (_HoneyGenerated *HoneyGeneratedTransactorSession) GetExchangable() (*types.Transaction, error) {
	return _HoneyGenerated.Contract.GetExchangable(&_HoneyGenerated.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xc6c3bbe6.
//
// Solidity: function mint(address to, address collateral, uint256 amount) returns(uint256)
func (_HoneyGenerated *HoneyGeneratedTransactor) Mint(opts *bind.TransactOpts, to common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoneyGenerated.contract.Transact(opts, "mint", to, collateral, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xc6c3bbe6.
//
// Solidity: function mint(address to, address collateral, uint256 amount) returns(uint256)
func (_HoneyGenerated *HoneyGeneratedSession) Mint(to common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoneyGenerated.Contract.Mint(&_HoneyGenerated.TransactOpts, to, collateral, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xc6c3bbe6.
//
// Solidity: function mint(address to, address collateral, uint256 amount) returns(uint256)
func (_HoneyGenerated *HoneyGeneratedTransactorSession) Mint(to common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoneyGenerated.Contract.Mint(&_HoneyGenerated.TransactOpts, to, collateral, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x5c833bfd.
//
// Solidity: function redeem(address to, uint256 amount, address collateral) returns(uint256)
func (_HoneyGenerated *HoneyGeneratedTransactor) Redeem(opts *bind.TransactOpts, to common.Address, amount *big.Int, collateral common.Address) (*types.Transaction, error) {
	return _HoneyGenerated.contract.Transact(opts, "redeem", to, amount, collateral)
}

// Redeem is a paid mutator transaction binding the contract method 0x5c833bfd.
//
// Solidity: function redeem(address to, uint256 amount, address collateral) returns(uint256)
func (_HoneyGenerated *HoneyGeneratedSession) Redeem(to common.Address, amount *big.Int, collateral common.Address) (*types.Transaction, error) {
	return _HoneyGenerated.Contract.Redeem(&_HoneyGenerated.TransactOpts, to, amount, collateral)
}

// Redeem is a paid mutator transaction binding the contract method 0x5c833bfd.
//
// Solidity: function redeem(address to, uint256 amount, address collateral) returns(uint256)
func (_HoneyGenerated *HoneyGeneratedTransactorSession) Redeem(to common.Address, amount *big.Int, collateral common.Address) (*types.Transaction, error) {
	return _HoneyGenerated.Contract.Redeem(&_HoneyGenerated.TransactOpts, to, amount, collateral)
}
