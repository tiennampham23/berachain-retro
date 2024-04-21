// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bera404

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

// Bera404MetaData contains all meta data concerning the Bera404 contract.
var Bera404MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRemoveFromWhitelist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DecimalsTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidApproval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedIndexOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitDeadlineExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeRecipient\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ERC20Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ERC721Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ERC721Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueOrId_\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"mintNFT\",\"type\":\"bool\"}],\"name\":\"boobaMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"erc20BalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20TotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"erc721BalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc721TokensBankedInQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc721TotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"owned\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"erc721Owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved_\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value_\",\"type\":\"bool\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueOrId_\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"units\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Bera404ABI is the input ABI used to generate the binding from.
// Deprecated: Use Bera404MetaData.ABI instead.
var Bera404ABI = Bera404MetaData.ABI

// Bera404 is an auto generated Go binding around an Ethereum contract.
type Bera404 struct {
	Bera404Caller     // Read-only binding to the contract
	Bera404Transactor // Write-only binding to the contract
	Bera404Filterer   // Log filterer for contract events
}

// Bera404Caller is an auto generated read-only Go binding around an Ethereum contract.
type Bera404Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bera404Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Bera404Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bera404Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bera404Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bera404Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bera404Session struct {
	Contract     *Bera404          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Bera404CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bera404CallerSession struct {
	Contract *Bera404Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Bera404TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bera404TransactorSession struct {
	Contract     *Bera404Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Bera404Raw is an auto generated low-level Go binding around an Ethereum contract.
type Bera404Raw struct {
	Contract *Bera404 // Generic contract binding to access the raw methods on
}

// Bera404CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bera404CallerRaw struct {
	Contract *Bera404Caller // Generic read-only contract binding to access the raw methods on
}

// Bera404TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bera404TransactorRaw struct {
	Contract *Bera404Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBera404 creates a new instance of Bera404, bound to a specific deployed contract.
func NewBera404(address common.Address, backend bind.ContractBackend) (*Bera404, error) {
	contract, err := bindBera404(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bera404{Bera404Caller: Bera404Caller{contract: contract}, Bera404Transactor: Bera404Transactor{contract: contract}, Bera404Filterer: Bera404Filterer{contract: contract}}, nil
}

// NewBera404Caller creates a new read-only instance of Bera404, bound to a specific deployed contract.
func NewBera404Caller(address common.Address, caller bind.ContractCaller) (*Bera404Caller, error) {
	contract, err := bindBera404(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bera404Caller{contract: contract}, nil
}

// NewBera404Transactor creates a new write-only instance of Bera404, bound to a specific deployed contract.
func NewBera404Transactor(address common.Address, transactor bind.ContractTransactor) (*Bera404Transactor, error) {
	contract, err := bindBera404(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bera404Transactor{contract: contract}, nil
}

// NewBera404Filterer creates a new log filterer instance of Bera404, bound to a specific deployed contract.
func NewBera404Filterer(address common.Address, filterer bind.ContractFilterer) (*Bera404Filterer, error) {
	contract, err := bindBera404(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bera404Filterer{contract: contract}, nil
}

// bindBera404 binds a generic wrapper to an already deployed contract.
func bindBera404(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Bera404MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bera404 *Bera404Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bera404.Contract.Bera404Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bera404 *Bera404Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bera404.Contract.Bera404Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bera404 *Bera404Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bera404.Contract.Bera404Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bera404 *Bera404CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bera404.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bera404 *Bera404TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bera404.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bera404 *Bera404TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bera404.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bera404 *Bera404Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bera404 *Bera404Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bera404.Contract.DOMAINSEPARATOR(&_Bera404.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bera404 *Bera404CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bera404.Contract.DOMAINSEPARATOR(&_Bera404.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Bera404 *Bera404Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Bera404 *Bera404Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Bera404.Contract.Allowance(&_Bera404.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Bera404 *Bera404CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Bera404.Contract.Allowance(&_Bera404.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Bera404 *Bera404Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Bera404 *Bera404Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Bera404.Contract.BalanceOf(&_Bera404.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Bera404 *Bera404CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Bera404.Contract.BalanceOf(&_Bera404.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bera404 *Bera404Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bera404 *Bera404Session) Decimals() (uint8, error) {
	return _Bera404.Contract.Decimals(&_Bera404.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bera404 *Bera404CallerSession) Decimals() (uint8, error) {
	return _Bera404.Contract.Decimals(&_Bera404.CallOpts)
}

// Erc20BalanceOf is a free data retrieval call binding the contract method 0x02519da3.
//
// Solidity: function erc20BalanceOf(address owner_) view returns(uint256)
func (_Bera404 *Bera404Caller) Erc20BalanceOf(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "erc20BalanceOf", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc20BalanceOf is a free data retrieval call binding the contract method 0x02519da3.
//
// Solidity: function erc20BalanceOf(address owner_) view returns(uint256)
func (_Bera404 *Bera404Session) Erc20BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _Bera404.Contract.Erc20BalanceOf(&_Bera404.CallOpts, owner_)
}

// Erc20BalanceOf is a free data retrieval call binding the contract method 0x02519da3.
//
// Solidity: function erc20BalanceOf(address owner_) view returns(uint256)
func (_Bera404 *Bera404CallerSession) Erc20BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _Bera404.Contract.Erc20BalanceOf(&_Bera404.CallOpts, owner_)
}

// Erc20TotalSupply is a free data retrieval call binding the contract method 0x89fb4c66.
//
// Solidity: function erc20TotalSupply() view returns(uint256)
func (_Bera404 *Bera404Caller) Erc20TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "erc20TotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc20TotalSupply is a free data retrieval call binding the contract method 0x89fb4c66.
//
// Solidity: function erc20TotalSupply() view returns(uint256)
func (_Bera404 *Bera404Session) Erc20TotalSupply() (*big.Int, error) {
	return _Bera404.Contract.Erc20TotalSupply(&_Bera404.CallOpts)
}

// Erc20TotalSupply is a free data retrieval call binding the contract method 0x89fb4c66.
//
// Solidity: function erc20TotalSupply() view returns(uint256)
func (_Bera404 *Bera404CallerSession) Erc20TotalSupply() (*big.Int, error) {
	return _Bera404.Contract.Erc20TotalSupply(&_Bera404.CallOpts)
}

// Erc721BalanceOf is a free data retrieval call binding the contract method 0xb3f9ea34.
//
// Solidity: function erc721BalanceOf(address owner_) view returns(uint256)
func (_Bera404 *Bera404Caller) Erc721BalanceOf(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "erc721BalanceOf", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc721BalanceOf is a free data retrieval call binding the contract method 0xb3f9ea34.
//
// Solidity: function erc721BalanceOf(address owner_) view returns(uint256)
func (_Bera404 *Bera404Session) Erc721BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _Bera404.Contract.Erc721BalanceOf(&_Bera404.CallOpts, owner_)
}

// Erc721BalanceOf is a free data retrieval call binding the contract method 0xb3f9ea34.
//
// Solidity: function erc721BalanceOf(address owner_) view returns(uint256)
func (_Bera404 *Bera404CallerSession) Erc721BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _Bera404.Contract.Erc721BalanceOf(&_Bera404.CallOpts, owner_)
}

// Erc721TokensBankedInQueue is a free data retrieval call binding the contract method 0x09d890d5.
//
// Solidity: function erc721TokensBankedInQueue() view returns(uint256)
func (_Bera404 *Bera404Caller) Erc721TokensBankedInQueue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "erc721TokensBankedInQueue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc721TokensBankedInQueue is a free data retrieval call binding the contract method 0x09d890d5.
//
// Solidity: function erc721TokensBankedInQueue() view returns(uint256)
func (_Bera404 *Bera404Session) Erc721TokensBankedInQueue() (*big.Int, error) {
	return _Bera404.Contract.Erc721TokensBankedInQueue(&_Bera404.CallOpts)
}

// Erc721TokensBankedInQueue is a free data retrieval call binding the contract method 0x09d890d5.
//
// Solidity: function erc721TokensBankedInQueue() view returns(uint256)
func (_Bera404 *Bera404CallerSession) Erc721TokensBankedInQueue() (*big.Int, error) {
	return _Bera404.Contract.Erc721TokensBankedInQueue(&_Bera404.CallOpts)
}

// Erc721TotalSupply is a free data retrieval call binding the contract method 0xc5ab3ba6.
//
// Solidity: function erc721TotalSupply() view returns(uint256)
func (_Bera404 *Bera404Caller) Erc721TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "erc721TotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc721TotalSupply is a free data retrieval call binding the contract method 0xc5ab3ba6.
//
// Solidity: function erc721TotalSupply() view returns(uint256)
func (_Bera404 *Bera404Session) Erc721TotalSupply() (*big.Int, error) {
	return _Bera404.Contract.Erc721TotalSupply(&_Bera404.CallOpts)
}

// Erc721TotalSupply is a free data retrieval call binding the contract method 0xc5ab3ba6.
//
// Solidity: function erc721TotalSupply() view returns(uint256)
func (_Bera404 *Bera404CallerSession) Erc721TotalSupply() (*big.Int, error) {
	return _Bera404.Contract.Erc721TotalSupply(&_Bera404.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Bera404 *Bera404Caller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Bera404 *Bera404Session) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Bera404.Contract.GetApproved(&_Bera404.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Bera404 *Bera404CallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Bera404.Contract.GetApproved(&_Bera404.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Bera404 *Bera404Caller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Bera404 *Bera404Session) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Bera404.Contract.IsApprovedForAll(&_Bera404.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Bera404 *Bera404CallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Bera404.Contract.IsApprovedForAll(&_Bera404.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bera404 *Bera404Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bera404 *Bera404Session) Name() (string, error) {
	return _Bera404.Contract.Name(&_Bera404.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bera404 *Bera404CallerSession) Name() (string, error) {
	return _Bera404.Contract.Name(&_Bera404.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Bera404 *Bera404Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Bera404 *Bera404Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Bera404.Contract.Nonces(&_Bera404.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Bera404 *Bera404CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Bera404.Contract.Nonces(&_Bera404.CallOpts, arg0)
}

// Owned is a free data retrieval call binding the contract method 0xb1ab9317.
//
// Solidity: function owned(address owner_) view returns(uint256[])
func (_Bera404 *Bera404Caller) Owned(opts *bind.CallOpts, owner_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "owned", owner_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Owned is a free data retrieval call binding the contract method 0xb1ab9317.
//
// Solidity: function owned(address owner_) view returns(uint256[])
func (_Bera404 *Bera404Session) Owned(owner_ common.Address) ([]*big.Int, error) {
	return _Bera404.Contract.Owned(&_Bera404.CallOpts, owner_)
}

// Owned is a free data retrieval call binding the contract method 0xb1ab9317.
//
// Solidity: function owned(address owner_) view returns(uint256[])
func (_Bera404 *Bera404CallerSession) Owned(owner_ common.Address) ([]*big.Int, error) {
	return _Bera404.Contract.Owned(&_Bera404.CallOpts, owner_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bera404 *Bera404Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bera404 *Bera404Session) Owner() (common.Address, error) {
	return _Bera404.Contract.Owner(&_Bera404.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bera404 *Bera404CallerSession) Owner() (common.Address, error) {
	return _Bera404.Contract.Owner(&_Bera404.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id_) view returns(address erc721Owner)
func (_Bera404 *Bera404Caller) OwnerOf(opts *bind.CallOpts, id_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "ownerOf", id_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id_) view returns(address erc721Owner)
func (_Bera404 *Bera404Session) OwnerOf(id_ *big.Int) (common.Address, error) {
	return _Bera404.Contract.OwnerOf(&_Bera404.CallOpts, id_)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id_) view returns(address erc721Owner)
func (_Bera404 *Bera404CallerSession) OwnerOf(id_ *big.Int) (common.Address, error) {
	return _Bera404.Contract.OwnerOf(&_Bera404.CallOpts, id_)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bera404 *Bera404Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bera404 *Bera404Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bera404.Contract.SupportsInterface(&_Bera404.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bera404 *Bera404CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bera404.Contract.SupportsInterface(&_Bera404.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bera404 *Bera404Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bera404 *Bera404Session) Symbol() (string, error) {
	return _Bera404.Contract.Symbol(&_Bera404.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bera404 *Bera404CallerSession) Symbol() (string, error) {
	return _Bera404.Contract.Symbol(&_Bera404.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id_) pure returns(string)
func (_Bera404 *Bera404Caller) TokenURI(opts *bind.CallOpts, id_ *big.Int) (string, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "tokenURI", id_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id_) pure returns(string)
func (_Bera404 *Bera404Session) TokenURI(id_ *big.Int) (string, error) {
	return _Bera404.Contract.TokenURI(&_Bera404.CallOpts, id_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id_) pure returns(string)
func (_Bera404 *Bera404CallerSession) TokenURI(id_ *big.Int) (string, error) {
	return _Bera404.Contract.TokenURI(&_Bera404.CallOpts, id_)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bera404 *Bera404Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bera404 *Bera404Session) TotalSupply() (*big.Int, error) {
	return _Bera404.Contract.TotalSupply(&_Bera404.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bera404 *Bera404CallerSession) TotalSupply() (*big.Int, error) {
	return _Bera404.Contract.TotalSupply(&_Bera404.CallOpts)
}

// Units is a free data retrieval call binding the contract method 0x976a8435.
//
// Solidity: function units() view returns(uint256)
func (_Bera404 *Bera404Caller) Units(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "units")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Units is a free data retrieval call binding the contract method 0x976a8435.
//
// Solidity: function units() view returns(uint256)
func (_Bera404 *Bera404Session) Units() (*big.Int, error) {
	return _Bera404.Contract.Units(&_Bera404.CallOpts)
}

// Units is a free data retrieval call binding the contract method 0x976a8435.
//
// Solidity: function units() view returns(uint256)
func (_Bera404 *Bera404CallerSession) Units() (*big.Int, error) {
	return _Bera404.Contract.Units(&_Bera404.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Bera404 *Bera404Caller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bera404.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Bera404 *Bera404Session) Whitelist(arg0 common.Address) (bool, error) {
	return _Bera404.Contract.Whitelist(&_Bera404.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Bera404 *Bera404CallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Bera404.Contract.Whitelist(&_Bera404.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender_, uint256 valueOrId_) returns(bool)
func (_Bera404 *Bera404Transactor) Approve(opts *bind.TransactOpts, spender_ common.Address, valueOrId_ *big.Int) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "approve", spender_, valueOrId_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender_, uint256 valueOrId_) returns(bool)
func (_Bera404 *Bera404Session) Approve(spender_ common.Address, valueOrId_ *big.Int) (*types.Transaction, error) {
	return _Bera404.Contract.Approve(&_Bera404.TransactOpts, spender_, valueOrId_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender_, uint256 valueOrId_) returns(bool)
func (_Bera404 *Bera404TransactorSession) Approve(spender_ common.Address, valueOrId_ *big.Int) (*types.Transaction, error) {
	return _Bera404.Contract.Approve(&_Bera404.TransactOpts, spender_, valueOrId_)
}

// BoobaMint is a paid mutator transaction binding the contract method 0x1f78fa9f.
//
// Solidity: function boobaMint(uint256 value_, bool mintNFT) returns()
func (_Bera404 *Bera404Transactor) BoobaMint(opts *bind.TransactOpts, value_ *big.Int, mintNFT bool) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "boobaMint", value_, mintNFT)
}

// BoobaMint is a paid mutator transaction binding the contract method 0x1f78fa9f.
//
// Solidity: function boobaMint(uint256 value_, bool mintNFT) returns()
func (_Bera404 *Bera404Session) BoobaMint(value_ *big.Int, mintNFT bool) (*types.Transaction, error) {
	return _Bera404.Contract.BoobaMint(&_Bera404.TransactOpts, value_, mintNFT)
}

// BoobaMint is a paid mutator transaction binding the contract method 0x1f78fa9f.
//
// Solidity: function boobaMint(uint256 value_, bool mintNFT) returns()
func (_Bera404 *Bera404TransactorSession) BoobaMint(value_ *big.Int, mintNFT bool) (*types.Transaction, error) {
	return _Bera404.Contract.BoobaMint(&_Bera404.TransactOpts, value_, mintNFT)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner_, address spender_, uint256 value_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns()
func (_Bera404 *Bera404Transactor) Permit(opts *bind.TransactOpts, owner_ common.Address, spender_ common.Address, value_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "permit", owner_, spender_, value_, deadline_, v_, r_, s_)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner_, address spender_, uint256 value_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns()
func (_Bera404 *Bera404Session) Permit(owner_ common.Address, spender_ common.Address, value_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _Bera404.Contract.Permit(&_Bera404.TransactOpts, owner_, spender_, value_, deadline_, v_, r_, s_)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner_, address spender_, uint256 value_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns()
func (_Bera404 *Bera404TransactorSession) Permit(owner_ common.Address, spender_ common.Address, value_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _Bera404.Contract.Permit(&_Bera404.TransactOpts, owner_, spender_, value_, deadline_, v_, r_, s_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bera404 *Bera404Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bera404 *Bera404Session) RenounceOwnership() (*types.Transaction, error) {
	return _Bera404.Contract.RenounceOwnership(&_Bera404.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bera404 *Bera404TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bera404.Contract.RenounceOwnership(&_Bera404.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 id_) returns()
func (_Bera404 *Bera404Transactor) SafeTransferFrom(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, id_ *big.Int) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "safeTransferFrom", from_, to_, id_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 id_) returns()
func (_Bera404 *Bera404Session) SafeTransferFrom(from_ common.Address, to_ common.Address, id_ *big.Int) (*types.Transaction, error) {
	return _Bera404.Contract.SafeTransferFrom(&_Bera404.TransactOpts, from_, to_, id_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 id_) returns()
func (_Bera404 *Bera404TransactorSession) SafeTransferFrom(from_ common.Address, to_ common.Address, id_ *big.Int) (*types.Transaction, error) {
	return _Bera404.Contract.SafeTransferFrom(&_Bera404.TransactOpts, from_, to_, id_)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 id_, bytes data_) returns()
func (_Bera404 *Bera404Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, id_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "safeTransferFrom0", from_, to_, id_, data_)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 id_, bytes data_) returns()
func (_Bera404 *Bera404Session) SafeTransferFrom0(from_ common.Address, to_ common.Address, id_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Bera404.Contract.SafeTransferFrom0(&_Bera404.TransactOpts, from_, to_, id_, data_)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 id_, bytes data_) returns()
func (_Bera404 *Bera404TransactorSession) SafeTransferFrom0(from_ common.Address, to_ common.Address, id_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Bera404.Contract.SafeTransferFrom0(&_Bera404.TransactOpts, from_, to_, id_, data_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved_) returns()
func (_Bera404 *Bera404Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator_ common.Address, approved_ bool) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "setApprovalForAll", operator_, approved_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved_) returns()
func (_Bera404 *Bera404Session) SetApprovalForAll(operator_ common.Address, approved_ bool) (*types.Transaction, error) {
	return _Bera404.Contract.SetApprovalForAll(&_Bera404.TransactOpts, operator_, approved_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved_) returns()
func (_Bera404 *Bera404TransactorSession) SetApprovalForAll(operator_ common.Address, approved_ bool) (*types.Transaction, error) {
	return _Bera404.Contract.SetApprovalForAll(&_Bera404.TransactOpts, operator_, approved_)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x53d6fd59.
//
// Solidity: function setWhitelist(address account_, bool value_) returns()
func (_Bera404 *Bera404Transactor) SetWhitelist(opts *bind.TransactOpts, account_ common.Address, value_ bool) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "setWhitelist", account_, value_)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x53d6fd59.
//
// Solidity: function setWhitelist(address account_, bool value_) returns()
func (_Bera404 *Bera404Session) SetWhitelist(account_ common.Address, value_ bool) (*types.Transaction, error) {
	return _Bera404.Contract.SetWhitelist(&_Bera404.TransactOpts, account_, value_)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x53d6fd59.
//
// Solidity: function setWhitelist(address account_, bool value_) returns()
func (_Bera404 *Bera404TransactorSession) SetWhitelist(account_ common.Address, value_ bool) (*types.Transaction, error) {
	return _Bera404.Contract.SetWhitelist(&_Bera404.TransactOpts, account_, value_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to_, uint256 value_) returns(bool)
func (_Bera404 *Bera404Transactor) Transfer(opts *bind.TransactOpts, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "transfer", to_, value_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to_, uint256 value_) returns(bool)
func (_Bera404 *Bera404Session) Transfer(to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Bera404.Contract.Transfer(&_Bera404.TransactOpts, to_, value_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to_, uint256 value_) returns(bool)
func (_Bera404 *Bera404TransactorSession) Transfer(to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Bera404.Contract.Transfer(&_Bera404.TransactOpts, to_, value_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 valueOrId_) returns(bool)
func (_Bera404 *Bera404Transactor) TransferFrom(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, valueOrId_ *big.Int) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "transferFrom", from_, to_, valueOrId_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 valueOrId_) returns(bool)
func (_Bera404 *Bera404Session) TransferFrom(from_ common.Address, to_ common.Address, valueOrId_ *big.Int) (*types.Transaction, error) {
	return _Bera404.Contract.TransferFrom(&_Bera404.TransactOpts, from_, to_, valueOrId_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 valueOrId_) returns(bool)
func (_Bera404 *Bera404TransactorSession) TransferFrom(from_ common.Address, to_ common.Address, valueOrId_ *big.Int) (*types.Transaction, error) {
	return _Bera404.Contract.TransferFrom(&_Bera404.TransactOpts, from_, to_, valueOrId_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bera404 *Bera404Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bera404.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bera404 *Bera404Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bera404.Contract.TransferOwnership(&_Bera404.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bera404 *Bera404TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bera404.Contract.TransferOwnership(&_Bera404.TransactOpts, newOwner)
}

// Bera404ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Bera404 contract.
type Bera404ApprovalForAllIterator struct {
	Event *Bera404ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Bera404ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bera404ApprovalForAll)
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
		it.Event = new(Bera404ApprovalForAll)
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
func (it *Bera404ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bera404ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bera404ApprovalForAll represents a ApprovalForAll event raised by the Bera404 contract.
type Bera404ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bera404 *Bera404Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Bera404ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bera404.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Bera404ApprovalForAllIterator{contract: _Bera404.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bera404 *Bera404Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Bera404ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bera404.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bera404ApprovalForAll)
				if err := _Bera404.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Bera404 *Bera404Filterer) ParseApprovalForAll(log types.Log) (*Bera404ApprovalForAll, error) {
	event := new(Bera404ApprovalForAll)
	if err := _Bera404.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bera404ERC20ApprovalIterator is returned from FilterERC20Approval and is used to iterate over the raw logs and unpacked data for ERC20Approval events raised by the Bera404 contract.
type Bera404ERC20ApprovalIterator struct {
	Event *Bera404ERC20Approval // Event containing the contract specifics and raw log

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
func (it *Bera404ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bera404ERC20Approval)
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
		it.Event = new(Bera404ERC20Approval)
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
func (it *Bera404ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bera404ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bera404ERC20Approval represents a ERC20Approval event raised by the Bera404 contract.
type Bera404ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC20Approval is a free log retrieval operation binding the contract event 0x1f01303a1ce9329d9963e1937c201e23c5543a9e3536e9edead087aec7dc6d83.
//
// Solidity: event ERC20Approval(address owner, address spender, uint256 value)
func (_Bera404 *Bera404Filterer) FilterERC20Approval(opts *bind.FilterOpts) (*Bera404ERC20ApprovalIterator, error) {

	logs, sub, err := _Bera404.contract.FilterLogs(opts, "ERC20Approval")
	if err != nil {
		return nil, err
	}
	return &Bera404ERC20ApprovalIterator{contract: _Bera404.contract, event: "ERC20Approval", logs: logs, sub: sub}, nil
}

// WatchERC20Approval is a free log subscription operation binding the contract event 0x1f01303a1ce9329d9963e1937c201e23c5543a9e3536e9edead087aec7dc6d83.
//
// Solidity: event ERC20Approval(address owner, address spender, uint256 value)
func (_Bera404 *Bera404Filterer) WatchERC20Approval(opts *bind.WatchOpts, sink chan<- *Bera404ERC20Approval) (event.Subscription, error) {

	logs, sub, err := _Bera404.contract.WatchLogs(opts, "ERC20Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bera404ERC20Approval)
				if err := _Bera404.contract.UnpackLog(event, "ERC20Approval", log); err != nil {
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

// ParseERC20Approval is a log parse operation binding the contract event 0x1f01303a1ce9329d9963e1937c201e23c5543a9e3536e9edead087aec7dc6d83.
//
// Solidity: event ERC20Approval(address owner, address spender, uint256 value)
func (_Bera404 *Bera404Filterer) ParseERC20Approval(log types.Log) (*Bera404ERC20Approval, error) {
	event := new(Bera404ERC20Approval)
	if err := _Bera404.contract.UnpackLog(event, "ERC20Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bera404ERC20TransferIterator is returned from FilterERC20Transfer and is used to iterate over the raw logs and unpacked data for ERC20Transfer events raised by the Bera404 contract.
type Bera404ERC20TransferIterator struct {
	Event *Bera404ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *Bera404ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bera404ERC20Transfer)
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
		it.Event = new(Bera404ERC20Transfer)
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
func (it *Bera404ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bera404ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bera404ERC20Transfer represents a ERC20Transfer event raised by the Bera404 contract.
type Bera404ERC20Transfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20Transfer is a free log retrieval operation binding the contract event 0xe59fdd36d0d223c0c7d996db7ad796880f45e1936cb0bb7ac102e7082e031487.
//
// Solidity: event ERC20Transfer(address indexed from, address indexed to, uint256 amount)
func (_Bera404 *Bera404Filterer) FilterERC20Transfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Bera404ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bera404.contract.FilterLogs(opts, "ERC20Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Bera404ERC20TransferIterator{contract: _Bera404.contract, event: "ERC20Transfer", logs: logs, sub: sub}, nil
}

// WatchERC20Transfer is a free log subscription operation binding the contract event 0xe59fdd36d0d223c0c7d996db7ad796880f45e1936cb0bb7ac102e7082e031487.
//
// Solidity: event ERC20Transfer(address indexed from, address indexed to, uint256 amount)
func (_Bera404 *Bera404Filterer) WatchERC20Transfer(opts *bind.WatchOpts, sink chan<- *Bera404ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bera404.contract.WatchLogs(opts, "ERC20Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bera404ERC20Transfer)
				if err := _Bera404.contract.UnpackLog(event, "ERC20Transfer", log); err != nil {
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

// ParseERC20Transfer is a log parse operation binding the contract event 0xe59fdd36d0d223c0c7d996db7ad796880f45e1936cb0bb7ac102e7082e031487.
//
// Solidity: event ERC20Transfer(address indexed from, address indexed to, uint256 amount)
func (_Bera404 *Bera404Filterer) ParseERC20Transfer(log types.Log) (*Bera404ERC20Transfer, error) {
	event := new(Bera404ERC20Transfer)
	if err := _Bera404.contract.UnpackLog(event, "ERC20Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bera404ERC721ApprovalIterator is returned from FilterERC721Approval and is used to iterate over the raw logs and unpacked data for ERC721Approval events raised by the Bera404 contract.
type Bera404ERC721ApprovalIterator struct {
	Event *Bera404ERC721Approval // Event containing the contract specifics and raw log

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
func (it *Bera404ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bera404ERC721Approval)
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
		it.Event = new(Bera404ERC721Approval)
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
func (it *Bera404ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bera404ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bera404ERC721Approval represents a ERC721Approval event raised by the Bera404 contract.
type Bera404ERC721Approval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC721Approval is a free log retrieval operation binding the contract event 0x797365dabb18fa726ccbccbe18c6f24c34e3b0653f2e99ea873bd7b84763dde6.
//
// Solidity: event ERC721Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Bera404 *Bera404Filterer) FilterERC721Approval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*Bera404ERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bera404.contract.FilterLogs(opts, "ERC721Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &Bera404ERC721ApprovalIterator{contract: _Bera404.contract, event: "ERC721Approval", logs: logs, sub: sub}, nil
}

// WatchERC721Approval is a free log subscription operation binding the contract event 0x797365dabb18fa726ccbccbe18c6f24c34e3b0653f2e99ea873bd7b84763dde6.
//
// Solidity: event ERC721Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Bera404 *Bera404Filterer) WatchERC721Approval(opts *bind.WatchOpts, sink chan<- *Bera404ERC721Approval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bera404.contract.WatchLogs(opts, "ERC721Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bera404ERC721Approval)
				if err := _Bera404.contract.UnpackLog(event, "ERC721Approval", log); err != nil {
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

// ParseERC721Approval is a log parse operation binding the contract event 0x797365dabb18fa726ccbccbe18c6f24c34e3b0653f2e99ea873bd7b84763dde6.
//
// Solidity: event ERC721Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Bera404 *Bera404Filterer) ParseERC721Approval(log types.Log) (*Bera404ERC721Approval, error) {
	event := new(Bera404ERC721Approval)
	if err := _Bera404.contract.UnpackLog(event, "ERC721Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bera404ERC721TransferIterator is returned from FilterERC721Transfer and is used to iterate over the raw logs and unpacked data for ERC721Transfer events raised by the Bera404 contract.
type Bera404ERC721TransferIterator struct {
	Event *Bera404ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *Bera404ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bera404ERC721Transfer)
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
		it.Event = new(Bera404ERC721Transfer)
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
func (it *Bera404ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bera404ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bera404ERC721Transfer represents a ERC721Transfer event raised by the Bera404 contract.
type Bera404ERC721Transfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterERC721Transfer is a free log retrieval operation binding the contract event 0xe5f815dc84b8cecdfd4beedfc3f91ab5be7af100eca4e8fb11552b867995394f.
//
// Solidity: event ERC721Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Bera404 *Bera404Filterer) FilterERC721Transfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*Bera404ERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bera404.contract.FilterLogs(opts, "ERC721Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &Bera404ERC721TransferIterator{contract: _Bera404.contract, event: "ERC721Transfer", logs: logs, sub: sub}, nil
}

// WatchERC721Transfer is a free log subscription operation binding the contract event 0xe5f815dc84b8cecdfd4beedfc3f91ab5be7af100eca4e8fb11552b867995394f.
//
// Solidity: event ERC721Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Bera404 *Bera404Filterer) WatchERC721Transfer(opts *bind.WatchOpts, sink chan<- *Bera404ERC721Transfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bera404.contract.WatchLogs(opts, "ERC721Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bera404ERC721Transfer)
				if err := _Bera404.contract.UnpackLog(event, "ERC721Transfer", log); err != nil {
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

// ParseERC721Transfer is a log parse operation binding the contract event 0xe5f815dc84b8cecdfd4beedfc3f91ab5be7af100eca4e8fb11552b867995394f.
//
// Solidity: event ERC721Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Bera404 *Bera404Filterer) ParseERC721Transfer(log types.Log) (*Bera404ERC721Transfer, error) {
	event := new(Bera404ERC721Transfer)
	if err := _Bera404.contract.UnpackLog(event, "ERC721Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bera404OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bera404 contract.
type Bera404OwnershipTransferredIterator struct {
	Event *Bera404OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Bera404OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bera404OwnershipTransferred)
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
		it.Event = new(Bera404OwnershipTransferred)
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
func (it *Bera404OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bera404OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bera404OwnershipTransferred represents a OwnershipTransferred event raised by the Bera404 contract.
type Bera404OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bera404 *Bera404Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Bera404OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bera404.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Bera404OwnershipTransferredIterator{contract: _Bera404.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bera404 *Bera404Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Bera404OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bera404.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bera404OwnershipTransferred)
				if err := _Bera404.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bera404 *Bera404Filterer) ParseOwnershipTransferred(log types.Log) (*Bera404OwnershipTransferred, error) {
	event := new(Bera404OwnershipTransferred)
	if err := _Bera404.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
