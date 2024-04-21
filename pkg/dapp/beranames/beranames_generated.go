// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package beranames

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

// BeranameGeneratedMetaData contains all meta data concerning the BeranameGenerated contract.
var BeranameGeneratedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAddressesProvider\",\"name\":\"addressesProvider_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Exists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeaseTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoEntity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Nope\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"chars\",\"type\":\"string[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"internalType\":\"contractIAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"chars\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundsManager\",\"outputs\":[{\"internalType\":\"contractIFundsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"chars\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"whois\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentAsset\",\"type\":\"address\"}],\"name\":\"mintERC20\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"chars\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"whois\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"singleEmojis\",\"type\":\"string[][]\"}],\"name\":\"mintToAuctionHouse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"minted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"whois\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"contractIPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"chars\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentAsset\",\"type\":\"address\"}],\"name\":\"renewERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"chars\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renewNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI_\",\"type\":\"string\"}],\"name\":\"updateMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"aka\",\"type\":\"address\"}],\"name\":\"updateWhois\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BeranameGeneratedABI is the input ABI used to generate the binding from.
// Deprecated: Use BeranameGeneratedMetaData.ABI instead.
var BeranameGeneratedABI = BeranameGeneratedMetaData.ABI

// BeranameGenerated is an auto generated Go binding around an Ethereum contract.
type BeranameGenerated struct {
	BeranameGeneratedCaller     // Read-only binding to the contract
	BeranameGeneratedTransactor // Write-only binding to the contract
	BeranameGeneratedFilterer   // Log filterer for contract events
}

// BeranameGeneratedCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeranameGeneratedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeranameGeneratedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeranameGeneratedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeranameGeneratedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeranameGeneratedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeranameGeneratedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeranameGeneratedSession struct {
	Contract     *BeranameGenerated // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BeranameGeneratedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeranameGeneratedCallerSession struct {
	Contract *BeranameGeneratedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BeranameGeneratedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeranameGeneratedTransactorSession struct {
	Contract     *BeranameGeneratedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BeranameGeneratedRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeranameGeneratedRaw struct {
	Contract *BeranameGenerated // Generic contract binding to access the raw methods on
}

// BeranameGeneratedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeranameGeneratedCallerRaw struct {
	Contract *BeranameGeneratedCaller // Generic read-only contract binding to access the raw methods on
}

// BeranameGeneratedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeranameGeneratedTransactorRaw struct {
	Contract *BeranameGeneratedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeranameGenerated creates a new instance of BeranameGenerated, bound to a specific deployed contract.
func NewBeranameGenerated(address common.Address, backend bind.ContractBackend) (*BeranameGenerated, error) {
	contract, err := bindBeranameGenerated(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeranameGenerated{BeranameGeneratedCaller: BeranameGeneratedCaller{contract: contract}, BeranameGeneratedTransactor: BeranameGeneratedTransactor{contract: contract}, BeranameGeneratedFilterer: BeranameGeneratedFilterer{contract: contract}}, nil
}

// NewBeranameGeneratedCaller creates a new read-only instance of BeranameGenerated, bound to a specific deployed contract.
func NewBeranameGeneratedCaller(address common.Address, caller bind.ContractCaller) (*BeranameGeneratedCaller, error) {
	contract, err := bindBeranameGenerated(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedCaller{contract: contract}, nil
}

// NewBeranameGeneratedTransactor creates a new write-only instance of BeranameGenerated, bound to a specific deployed contract.
func NewBeranameGeneratedTransactor(address common.Address, transactor bind.ContractTransactor) (*BeranameGeneratedTransactor, error) {
	contract, err := bindBeranameGenerated(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedTransactor{contract: contract}, nil
}

// NewBeranameGeneratedFilterer creates a new log filterer instance of BeranameGenerated, bound to a specific deployed contract.
func NewBeranameGeneratedFilterer(address common.Address, filterer bind.ContractFilterer) (*BeranameGeneratedFilterer, error) {
	contract, err := bindBeranameGenerated(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedFilterer{contract: contract}, nil
}

// bindBeranameGenerated binds a generic wrapper to an already deployed contract.
func bindBeranameGenerated(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeranameGeneratedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeranameGenerated *BeranameGeneratedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeranameGenerated.Contract.BeranameGeneratedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeranameGenerated *BeranameGeneratedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.BeranameGeneratedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeranameGenerated *BeranameGeneratedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.BeranameGeneratedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeranameGenerated *BeranameGeneratedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeranameGenerated.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeranameGenerated *BeranameGeneratedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeranameGenerated *BeranameGeneratedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.contract.Transact(opts, method, params...)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCaller) AddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "addressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_BeranameGenerated *BeranameGeneratedSession) AddressesProvider() (common.Address, error) {
	return _BeranameGenerated.Contract.AddressesProvider(&_BeranameGenerated.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCallerSession) AddressesProvider() (common.Address, error) {
	return _BeranameGenerated.Contract.AddressesProvider(&_BeranameGenerated.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BeranameGenerated.Contract.BalanceOf(&_BeranameGenerated.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BeranameGenerated.Contract.BalanceOf(&_BeranameGenerated.CallOpts, owner)
}

// Chars is a free data retrieval call binding the contract method 0xea372097.
//
// Solidity: function chars(uint256 id) view returns(string[])
func (_BeranameGenerated *BeranameGeneratedCaller) Chars(opts *bind.CallOpts, id *big.Int) ([]string, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "chars", id)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// Chars is a free data retrieval call binding the contract method 0xea372097.
//
// Solidity: function chars(uint256 id) view returns(string[])
func (_BeranameGenerated *BeranameGeneratedSession) Chars(id *big.Int) ([]string, error) {
	return _BeranameGenerated.Contract.Chars(&_BeranameGenerated.CallOpts, id)
}

// Chars is a free data retrieval call binding the contract method 0xea372097.
//
// Solidity: function chars(uint256 id) view returns(string[])
func (_BeranameGenerated *BeranameGeneratedCallerSession) Chars(id *big.Int) ([]string, error) {
	return _BeranameGenerated.Contract.Chars(&_BeranameGenerated.CallOpts, id)
}

// FundsManager is a free data retrieval call binding the contract method 0x0d116652.
//
// Solidity: function fundsManager() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCaller) FundsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "fundsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundsManager is a free data retrieval call binding the contract method 0x0d116652.
//
// Solidity: function fundsManager() view returns(address)
func (_BeranameGenerated *BeranameGeneratedSession) FundsManager() (common.Address, error) {
	return _BeranameGenerated.Contract.FundsManager(&_BeranameGenerated.CallOpts)
}

// FundsManager is a free data retrieval call binding the contract method 0x0d116652.
//
// Solidity: function fundsManager() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCallerSession) FundsManager() (common.Address, error) {
	return _BeranameGenerated.Contract.FundsManager(&_BeranameGenerated.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BeranameGenerated *BeranameGeneratedCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BeranameGenerated *BeranameGeneratedSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BeranameGenerated.Contract.GetApproved(&_BeranameGenerated.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BeranameGenerated *BeranameGeneratedCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BeranameGenerated.Contract.GetApproved(&_BeranameGenerated.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BeranameGenerated *BeranameGeneratedSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BeranameGenerated.Contract.IsApprovedForAll(&_BeranameGenerated.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BeranameGenerated.Contract.IsApprovedForAll(&_BeranameGenerated.CallOpts, owner, operator)
}

// Minted is a free data retrieval call binding the contract method 0x8ccc5f80.
//
// Solidity: function minted(bytes32 ) view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCaller) Minted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "minted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x8ccc5f80.
//
// Solidity: function minted(bytes32 ) view returns(bool)
func (_BeranameGenerated *BeranameGeneratedSession) Minted(arg0 [32]byte) (bool, error) {
	return _BeranameGenerated.Contract.Minted(&_BeranameGenerated.CallOpts, arg0)
}

// Minted is a free data retrieval call binding the contract method 0x8ccc5f80.
//
// Solidity: function minted(bytes32 ) view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCallerSession) Minted(arg0 [32]byte) (bool, error) {
	return _BeranameGenerated.Contract.Minted(&_BeranameGenerated.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeranameGenerated *BeranameGeneratedCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeranameGenerated *BeranameGeneratedSession) Name() (string, error) {
	return _BeranameGenerated.Contract.Name(&_BeranameGenerated.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeranameGenerated *BeranameGeneratedCallerSession) Name() (string, error) {
	return _BeranameGenerated.Contract.Name(&_BeranameGenerated.CallOpts)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(bytes32 name, uint256 expiry, address whois, string metadataURI)
func (_BeranameGenerated *BeranameGeneratedCaller) Names(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name        [32]byte
	Expiry      *big.Int
	Whois       common.Address
	MetadataURI string
}, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "names", arg0)

	outstruct := new(struct {
		Name        [32]byte
		Expiry      *big.Int
		Whois       common.Address
		MetadataURI string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Whois = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MetadataURI = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(bytes32 name, uint256 expiry, address whois, string metadataURI)
func (_BeranameGenerated *BeranameGeneratedSession) Names(arg0 *big.Int) (struct {
	Name        [32]byte
	Expiry      *big.Int
	Whois       common.Address
	MetadataURI string
}, error) {
	return _BeranameGenerated.Contract.Names(&_BeranameGenerated.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(bytes32 name, uint256 expiry, address whois, string metadataURI)
func (_BeranameGenerated *BeranameGeneratedCallerSession) Names(arg0 *big.Int) (struct {
	Name        [32]byte
	Expiry      *big.Int
	Whois       common.Address
	MetadataURI string
}, error) {
	return _BeranameGenerated.Contract.Names(&_BeranameGenerated.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeranameGenerated *BeranameGeneratedSession) Owner() (common.Address, error) {
	return _BeranameGenerated.Contract.Owner(&_BeranameGenerated.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCallerSession) Owner() (common.Address, error) {
	return _BeranameGenerated.Contract.Owner(&_BeranameGenerated.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BeranameGenerated *BeranameGeneratedCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BeranameGenerated *BeranameGeneratedSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BeranameGenerated.Contract.OwnerOf(&_BeranameGenerated.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BeranameGenerated *BeranameGeneratedCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BeranameGenerated.Contract.OwnerOf(&_BeranameGenerated.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BeranameGenerated *BeranameGeneratedSession) Paused() (bool, error) {
	return _BeranameGenerated.Contract.Paused(&_BeranameGenerated.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCallerSession) Paused() (bool, error) {
	return _BeranameGenerated.Contract.Paused(&_BeranameGenerated.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BeranameGenerated *BeranameGeneratedSession) PendingOwner() (common.Address, error) {
	return _BeranameGenerated.Contract.PendingOwner(&_BeranameGenerated.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCallerSession) PendingOwner() (common.Address, error) {
	return _BeranameGenerated.Contract.PendingOwner(&_BeranameGenerated.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_BeranameGenerated *BeranameGeneratedSession) PriceOracle() (common.Address, error) {
	return _BeranameGenerated.Contract.PriceOracle(&_BeranameGenerated.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_BeranameGenerated *BeranameGeneratedCallerSession) PriceOracle() (common.Address, error) {
	return _BeranameGenerated.Contract.PriceOracle(&_BeranameGenerated.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BeranameGenerated *BeranameGeneratedSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BeranameGenerated.Contract.SupportsInterface(&_BeranameGenerated.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BeranameGenerated.Contract.SupportsInterface(&_BeranameGenerated.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeranameGenerated *BeranameGeneratedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeranameGenerated *BeranameGeneratedSession) Symbol() (string, error) {
	return _BeranameGenerated.Contract.Symbol(&_BeranameGenerated.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeranameGenerated *BeranameGeneratedCallerSession) Symbol() (string, error) {
	return _BeranameGenerated.Contract.Symbol(&_BeranameGenerated.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _BeranameGenerated.Contract.TokenByIndex(&_BeranameGenerated.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _BeranameGenerated.Contract.TokenByIndex(&_BeranameGenerated.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _BeranameGenerated.Contract.TokenOfOwnerByIndex(&_BeranameGenerated.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _BeranameGenerated.Contract.TokenOfOwnerByIndex(&_BeranameGenerated.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_BeranameGenerated *BeranameGeneratedCaller) TokenURI(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "tokenURI", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_BeranameGenerated *BeranameGeneratedSession) TokenURI(id *big.Int) (string, error) {
	return _BeranameGenerated.Contract.TokenURI(&_BeranameGenerated.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_BeranameGenerated *BeranameGeneratedCallerSession) TokenURI(id *big.Int) (string, error) {
	return _BeranameGenerated.Contract.TokenURI(&_BeranameGenerated.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedSession) TotalSupply() (*big.Int, error) {
	return _BeranameGenerated.Contract.TotalSupply(&_BeranameGenerated.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeranameGenerated *BeranameGeneratedCallerSession) TotalSupply() (*big.Int, error) {
	return _BeranameGenerated.Contract.TotalSupply(&_BeranameGenerated.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BeranameGenerated.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_BeranameGenerated *BeranameGeneratedSession) WhitelistEnabled() (bool, error) {
	return _BeranameGenerated.Contract.WhitelistEnabled(&_BeranameGenerated.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_BeranameGenerated *BeranameGeneratedCallerSession) WhitelistEnabled() (bool, error) {
	return _BeranameGenerated.Contract.WhitelistEnabled(&_BeranameGenerated.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BeranameGenerated *BeranameGeneratedSession) AcceptOwnership() (*types.Transaction, error) {
	return _BeranameGenerated.Contract.AcceptOwnership(&_BeranameGenerated.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BeranameGenerated.Contract.AcceptOwnership(&_BeranameGenerated.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BeranameGenerated *BeranameGeneratedSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.Approve(&_BeranameGenerated.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.Approve(&_BeranameGenerated.TransactOpts, to, tokenId)
}

// MintERC20 is a paid mutator transaction binding the contract method 0xadca948d.
//
// Solidity: function mintERC20(string[] chars, uint256 duration, address whois, string metadataURI, address to, address paymentAsset) returns(uint256)
func (_BeranameGenerated *BeranameGeneratedTransactor) MintERC20(opts *bind.TransactOpts, chars []string, duration *big.Int, whois common.Address, metadataURI string, to common.Address, paymentAsset common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "mintERC20", chars, duration, whois, metadataURI, to, paymentAsset)
}

// MintERC20 is a paid mutator transaction binding the contract method 0xadca948d.
//
// Solidity: function mintERC20(string[] chars, uint256 duration, address whois, string metadataURI, address to, address paymentAsset) returns(uint256)
func (_BeranameGenerated *BeranameGeneratedSession) MintERC20(chars []string, duration *big.Int, whois common.Address, metadataURI string, to common.Address, paymentAsset common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.MintERC20(&_BeranameGenerated.TransactOpts, chars, duration, whois, metadataURI, to, paymentAsset)
}

// MintERC20 is a paid mutator transaction binding the contract method 0xadca948d.
//
// Solidity: function mintERC20(string[] chars, uint256 duration, address whois, string metadataURI, address to, address paymentAsset) returns(uint256)
func (_BeranameGenerated *BeranameGeneratedTransactorSession) MintERC20(chars []string, duration *big.Int, whois common.Address, metadataURI string, to common.Address, paymentAsset common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.MintERC20(&_BeranameGenerated.TransactOpts, chars, duration, whois, metadataURI, to, paymentAsset)
}

// MintNative is a paid mutator transaction binding the contract method 0x3d30c7f6.
//
// Solidity: function mintNative(string[] chars, uint256 duration, address whois, string metadataURI, address to) payable returns(uint256)
func (_BeranameGenerated *BeranameGeneratedTransactor) MintNative(opts *bind.TransactOpts, chars []string, duration *big.Int, whois common.Address, metadataURI string, to common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "mintNative", chars, duration, whois, metadataURI, to)
}

// MintNative is a paid mutator transaction binding the contract method 0x3d30c7f6.
//
// Solidity: function mintNative(string[] chars, uint256 duration, address whois, string metadataURI, address to) payable returns(uint256)
func (_BeranameGenerated *BeranameGeneratedSession) MintNative(chars []string, duration *big.Int, whois common.Address, metadataURI string, to common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.MintNative(&_BeranameGenerated.TransactOpts, chars, duration, whois, metadataURI, to)
}

// MintNative is a paid mutator transaction binding the contract method 0x3d30c7f6.
//
// Solidity: function mintNative(string[] chars, uint256 duration, address whois, string metadataURI, address to) payable returns(uint256)
func (_BeranameGenerated *BeranameGeneratedTransactorSession) MintNative(chars []string, duration *big.Int, whois common.Address, metadataURI string, to common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.MintNative(&_BeranameGenerated.TransactOpts, chars, duration, whois, metadataURI, to)
}

// MintToAuctionHouse is a paid mutator transaction binding the contract method 0x7b850843.
//
// Solidity: function mintToAuctionHouse(string[][] singleEmojis) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) MintToAuctionHouse(opts *bind.TransactOpts, singleEmojis [][]string) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "mintToAuctionHouse", singleEmojis)
}

// MintToAuctionHouse is a paid mutator transaction binding the contract method 0x7b850843.
//
// Solidity: function mintToAuctionHouse(string[][] singleEmojis) returns()
func (_BeranameGenerated *BeranameGeneratedSession) MintToAuctionHouse(singleEmojis [][]string) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.MintToAuctionHouse(&_BeranameGenerated.TransactOpts, singleEmojis)
}

// MintToAuctionHouse is a paid mutator transaction binding the contract method 0x7b850843.
//
// Solidity: function mintToAuctionHouse(string[][] singleEmojis) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) MintToAuctionHouse(singleEmojis [][]string) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.MintToAuctionHouse(&_BeranameGenerated.TransactOpts, singleEmojis)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BeranameGenerated *BeranameGeneratedTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BeranameGenerated *BeranameGeneratedSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.Multicall(&_BeranameGenerated.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BeranameGenerated *BeranameGeneratedTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.Multicall(&_BeranameGenerated.TransactOpts, data)
}

// RenewERC20 is a paid mutator transaction binding the contract method 0x3aa7ab05.
//
// Solidity: function renewERC20(string[] chars, uint256 duration, address paymentAsset) payable returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) RenewERC20(opts *bind.TransactOpts, chars []string, duration *big.Int, paymentAsset common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "renewERC20", chars, duration, paymentAsset)
}

// RenewERC20 is a paid mutator transaction binding the contract method 0x3aa7ab05.
//
// Solidity: function renewERC20(string[] chars, uint256 duration, address paymentAsset) payable returns()
func (_BeranameGenerated *BeranameGeneratedSession) RenewERC20(chars []string, duration *big.Int, paymentAsset common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.RenewERC20(&_BeranameGenerated.TransactOpts, chars, duration, paymentAsset)
}

// RenewERC20 is a paid mutator transaction binding the contract method 0x3aa7ab05.
//
// Solidity: function renewERC20(string[] chars, uint256 duration, address paymentAsset) payable returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) RenewERC20(chars []string, duration *big.Int, paymentAsset common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.RenewERC20(&_BeranameGenerated.TransactOpts, chars, duration, paymentAsset)
}

// RenewNative is a paid mutator transaction binding the contract method 0xe7962bb3.
//
// Solidity: function renewNative(string[] chars, uint256 duration) payable returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) RenewNative(opts *bind.TransactOpts, chars []string, duration *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "renewNative", chars, duration)
}

// RenewNative is a paid mutator transaction binding the contract method 0xe7962bb3.
//
// Solidity: function renewNative(string[] chars, uint256 duration) payable returns()
func (_BeranameGenerated *BeranameGeneratedSession) RenewNative(chars []string, duration *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.RenewNative(&_BeranameGenerated.TransactOpts, chars, duration)
}

// RenewNative is a paid mutator transaction binding the contract method 0xe7962bb3.
//
// Solidity: function renewNative(string[] chars, uint256 duration) payable returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) RenewNative(chars []string, duration *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.RenewNative(&_BeranameGenerated.TransactOpts, chars, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeranameGenerated *BeranameGeneratedSession) RenounceOwnership() (*types.Transaction, error) {
	return _BeranameGenerated.Contract.RenounceOwnership(&_BeranameGenerated.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BeranameGenerated.Contract.RenounceOwnership(&_BeranameGenerated.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BeranameGenerated *BeranameGeneratedSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.SafeTransferFrom(&_BeranameGenerated.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.SafeTransferFrom(&_BeranameGenerated.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BeranameGenerated *BeranameGeneratedSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.SafeTransferFrom0(&_BeranameGenerated.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.SafeTransferFrom0(&_BeranameGenerated.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BeranameGenerated *BeranameGeneratedSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.SetApprovalForAll(&_BeranameGenerated.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.SetApprovalForAll(&_BeranameGenerated.TransactOpts, operator, approved)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0xf3c4b704.
//
// Solidity: function setWhitelisted(address[] accounts, bool status) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) SetWhitelisted(opts *bind.TransactOpts, accounts []common.Address, status bool) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "setWhitelisted", accounts, status)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0xf3c4b704.
//
// Solidity: function setWhitelisted(address[] accounts, bool status) returns()
func (_BeranameGenerated *BeranameGeneratedSession) SetWhitelisted(accounts []common.Address, status bool) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.SetWhitelisted(&_BeranameGenerated.TransactOpts, accounts, status)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0xf3c4b704.
//
// Solidity: function setWhitelisted(address[] accounts, bool status) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) SetWhitelisted(accounts []common.Address, status bool) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.SetWhitelisted(&_BeranameGenerated.TransactOpts, accounts, status)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_BeranameGenerated *BeranameGeneratedSession) TogglePause() (*types.Transaction, error) {
	return _BeranameGenerated.Contract.TogglePause(&_BeranameGenerated.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) TogglePause() (*types.Transaction, error) {
	return _BeranameGenerated.Contract.TogglePause(&_BeranameGenerated.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BeranameGenerated *BeranameGeneratedSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.TransferFrom(&_BeranameGenerated.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.TransferFrom(&_BeranameGenerated.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeranameGenerated *BeranameGeneratedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.TransferOwnership(&_BeranameGenerated.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.TransferOwnership(&_BeranameGenerated.TransactOpts, newOwner)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x8fe86b3f.
//
// Solidity: function updateMetadataURI(uint256 id, string metadataURI_) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) UpdateMetadataURI(opts *bind.TransactOpts, id *big.Int, metadataURI_ string) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "updateMetadataURI", id, metadataURI_)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x8fe86b3f.
//
// Solidity: function updateMetadataURI(uint256 id, string metadataURI_) returns()
func (_BeranameGenerated *BeranameGeneratedSession) UpdateMetadataURI(id *big.Int, metadataURI_ string) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.UpdateMetadataURI(&_BeranameGenerated.TransactOpts, id, metadataURI_)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x8fe86b3f.
//
// Solidity: function updateMetadataURI(uint256 id, string metadataURI_) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) UpdateMetadataURI(id *big.Int, metadataURI_ string) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.UpdateMetadataURI(&_BeranameGenerated.TransactOpts, id, metadataURI_)
}

// UpdateWhois is a paid mutator transaction binding the contract method 0xbfb14ca7.
//
// Solidity: function updateWhois(uint256 id, address aka) returns()
func (_BeranameGenerated *BeranameGeneratedTransactor) UpdateWhois(opts *bind.TransactOpts, id *big.Int, aka common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.contract.Transact(opts, "updateWhois", id, aka)
}

// UpdateWhois is a paid mutator transaction binding the contract method 0xbfb14ca7.
//
// Solidity: function updateWhois(uint256 id, address aka) returns()
func (_BeranameGenerated *BeranameGeneratedSession) UpdateWhois(id *big.Int, aka common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.UpdateWhois(&_BeranameGenerated.TransactOpts, id, aka)
}

// UpdateWhois is a paid mutator transaction binding the contract method 0xbfb14ca7.
//
// Solidity: function updateWhois(uint256 id, address aka) returns()
func (_BeranameGenerated *BeranameGeneratedTransactorSession) UpdateWhois(id *big.Int, aka common.Address) (*types.Transaction, error) {
	return _BeranameGenerated.Contract.UpdateWhois(&_BeranameGenerated.TransactOpts, id, aka)
}

// BeranameGeneratedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BeranameGenerated contract.
type BeranameGeneratedApprovalIterator struct {
	Event *BeranameGeneratedApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BeranameGeneratedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeranameGeneratedApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BeranameGeneratedApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BeranameGeneratedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeranameGeneratedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeranameGeneratedApproval represents a Approval event raised by the BeranameGenerated contract.
type BeranameGeneratedApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BeranameGenerated *BeranameGeneratedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BeranameGeneratedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BeranameGenerated.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedApprovalIterator{contract: _BeranameGenerated.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BeranameGenerated *BeranameGeneratedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BeranameGeneratedApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BeranameGenerated.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeranameGeneratedApproval)
				if err := _BeranameGenerated.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BeranameGenerated *BeranameGeneratedFilterer) ParseApproval(log types.Log) (*BeranameGeneratedApproval, error) {
	event := new(BeranameGeneratedApproval)
	if err := _BeranameGenerated.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeranameGeneratedApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the BeranameGenerated contract.
type BeranameGeneratedApprovalForAllIterator struct {
	Event *BeranameGeneratedApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BeranameGeneratedApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeranameGeneratedApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BeranameGeneratedApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BeranameGeneratedApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeranameGeneratedApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeranameGeneratedApprovalForAll represents a ApprovalForAll event raised by the BeranameGenerated contract.
type BeranameGeneratedApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BeranameGenerated *BeranameGeneratedFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BeranameGeneratedApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BeranameGenerated.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedApprovalForAllIterator{contract: _BeranameGenerated.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BeranameGenerated *BeranameGeneratedFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BeranameGeneratedApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BeranameGenerated.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeranameGeneratedApprovalForAll)
				if err := _BeranameGenerated.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BeranameGenerated *BeranameGeneratedFilterer) ParseApprovalForAll(log types.Log) (*BeranameGeneratedApprovalForAll, error) {
	event := new(BeranameGeneratedApprovalForAll)
	if err := _BeranameGenerated.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeranameGeneratedMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BeranameGenerated contract.
type BeranameGeneratedMintIterator struct {
	Event *BeranameGeneratedMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BeranameGeneratedMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeranameGeneratedMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BeranameGeneratedMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BeranameGeneratedMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeranameGeneratedMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeranameGeneratedMint represents a Mint event raised by the BeranameGenerated contract.
type BeranameGeneratedMint struct {
	Id    *big.Int
	Chars []string
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xc368b9d544dedfe16e4ac071dcfbafe213adfe63bbc35e0fd438a2e815e98a44.
//
// Solidity: event Mint(uint256 indexed id, string[] chars, address indexed to)
func (_BeranameGenerated *BeranameGeneratedFilterer) FilterMint(opts *bind.FilterOpts, id []*big.Int, to []common.Address) (*BeranameGeneratedMintIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeranameGenerated.contract.FilterLogs(opts, "Mint", idRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedMintIterator{contract: _BeranameGenerated.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xc368b9d544dedfe16e4ac071dcfbafe213adfe63bbc35e0fd438a2e815e98a44.
//
// Solidity: event Mint(uint256 indexed id, string[] chars, address indexed to)
func (_BeranameGenerated *BeranameGeneratedFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BeranameGeneratedMint, id []*big.Int, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeranameGenerated.contract.WatchLogs(opts, "Mint", idRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeranameGeneratedMint)
				if err := _BeranameGenerated.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0xc368b9d544dedfe16e4ac071dcfbafe213adfe63bbc35e0fd438a2e815e98a44.
//
// Solidity: event Mint(uint256 indexed id, string[] chars, address indexed to)
func (_BeranameGenerated *BeranameGeneratedFilterer) ParseMint(log types.Log) (*BeranameGeneratedMint, error) {
	event := new(BeranameGeneratedMint)
	if err := _BeranameGenerated.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeranameGeneratedOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the BeranameGenerated contract.
type BeranameGeneratedOwnershipTransferStartedIterator struct {
	Event *BeranameGeneratedOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BeranameGeneratedOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeranameGeneratedOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BeranameGeneratedOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BeranameGeneratedOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeranameGeneratedOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeranameGeneratedOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the BeranameGenerated contract.
type BeranameGeneratedOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BeranameGenerated *BeranameGeneratedFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BeranameGeneratedOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeranameGenerated.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedOwnershipTransferStartedIterator{contract: _BeranameGenerated.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BeranameGenerated *BeranameGeneratedFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *BeranameGeneratedOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeranameGenerated.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeranameGeneratedOwnershipTransferStarted)
				if err := _BeranameGenerated.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BeranameGenerated *BeranameGeneratedFilterer) ParseOwnershipTransferStarted(log types.Log) (*BeranameGeneratedOwnershipTransferStarted, error) {
	event := new(BeranameGeneratedOwnershipTransferStarted)
	if err := _BeranameGenerated.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeranameGeneratedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BeranameGenerated contract.
type BeranameGeneratedOwnershipTransferredIterator struct {
	Event *BeranameGeneratedOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BeranameGeneratedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeranameGeneratedOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BeranameGeneratedOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BeranameGeneratedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeranameGeneratedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeranameGeneratedOwnershipTransferred represents a OwnershipTransferred event raised by the BeranameGenerated contract.
type BeranameGeneratedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BeranameGenerated *BeranameGeneratedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BeranameGeneratedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeranameGenerated.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedOwnershipTransferredIterator{contract: _BeranameGenerated.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BeranameGenerated *BeranameGeneratedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BeranameGeneratedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeranameGenerated.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeranameGeneratedOwnershipTransferred)
				if err := _BeranameGenerated.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BeranameGenerated *BeranameGeneratedFilterer) ParseOwnershipTransferred(log types.Log) (*BeranameGeneratedOwnershipTransferred, error) {
	event := new(BeranameGeneratedOwnershipTransferred)
	if err := _BeranameGenerated.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeranameGeneratedPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BeranameGenerated contract.
type BeranameGeneratedPausedIterator struct {
	Event *BeranameGeneratedPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BeranameGeneratedPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeranameGeneratedPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BeranameGeneratedPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BeranameGeneratedPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeranameGeneratedPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeranameGeneratedPaused represents a Paused event raised by the BeranameGenerated contract.
type BeranameGeneratedPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BeranameGenerated *BeranameGeneratedFilterer) FilterPaused(opts *bind.FilterOpts) (*BeranameGeneratedPausedIterator, error) {

	logs, sub, err := _BeranameGenerated.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedPausedIterator{contract: _BeranameGenerated.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BeranameGenerated *BeranameGeneratedFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BeranameGeneratedPaused) (event.Subscription, error) {

	logs, sub, err := _BeranameGenerated.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeranameGeneratedPaused)
				if err := _BeranameGenerated.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BeranameGenerated *BeranameGeneratedFilterer) ParsePaused(log types.Log) (*BeranameGeneratedPaused, error) {
	event := new(BeranameGeneratedPaused)
	if err := _BeranameGenerated.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeranameGeneratedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BeranameGenerated contract.
type BeranameGeneratedTransferIterator struct {
	Event *BeranameGeneratedTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BeranameGeneratedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeranameGeneratedTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BeranameGeneratedTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BeranameGeneratedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeranameGeneratedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeranameGeneratedTransfer represents a Transfer event raised by the BeranameGenerated contract.
type BeranameGeneratedTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BeranameGenerated *BeranameGeneratedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BeranameGeneratedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BeranameGenerated.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedTransferIterator{contract: _BeranameGenerated.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BeranameGenerated *BeranameGeneratedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BeranameGeneratedTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BeranameGenerated.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeranameGeneratedTransfer)
				if err := _BeranameGenerated.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BeranameGenerated *BeranameGeneratedFilterer) ParseTransfer(log types.Log) (*BeranameGeneratedTransfer, error) {
	event := new(BeranameGeneratedTransfer)
	if err := _BeranameGenerated.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeranameGeneratedUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BeranameGenerated contract.
type BeranameGeneratedUnpausedIterator struct {
	Event *BeranameGeneratedUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BeranameGeneratedUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeranameGeneratedUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BeranameGeneratedUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BeranameGeneratedUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeranameGeneratedUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeranameGeneratedUnpaused represents a Unpaused event raised by the BeranameGenerated contract.
type BeranameGeneratedUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BeranameGenerated *BeranameGeneratedFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BeranameGeneratedUnpausedIterator, error) {

	logs, sub, err := _BeranameGenerated.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BeranameGeneratedUnpausedIterator{contract: _BeranameGenerated.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BeranameGenerated *BeranameGeneratedFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BeranameGeneratedUnpaused) (event.Subscription, error) {

	logs, sub, err := _BeranameGenerated.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeranameGeneratedUnpaused)
				if err := _BeranameGenerated.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BeranameGenerated *BeranameGeneratedFilterer) ParseUnpaused(log types.Log) (*BeranameGeneratedUnpaused, error) {
	event := new(BeranameGeneratedUnpaused)
	if err := _BeranameGenerated.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
