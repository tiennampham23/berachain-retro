// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// Erc20GeneratedMetaData contains all meta data concerning the Erc20Generated contract.
var Erc20GeneratedMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}]},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"allowance\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"approve\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"balanceOf\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"decimals\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"type\":\"function\",\"name\":\"name\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"type\":\"function\",\"name\":\"symbol\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"type\":\"function\",\"name\":\"totalSupply\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"transfer\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"transferFrom\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]}]",
}

// Erc20GeneratedABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20GeneratedMetaData.ABI instead.
var Erc20GeneratedABI = Erc20GeneratedMetaData.ABI

// Erc20Generated is an auto generated Go binding around an Ethereum contract.
type Erc20Generated struct {
	Erc20GeneratedCaller     // Read-only binding to the contract
	Erc20GeneratedTransactor // Write-only binding to the contract
	Erc20GeneratedFilterer   // Log filterer for contract events
}

// Erc20GeneratedCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20GeneratedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20GeneratedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20GeneratedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20GeneratedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20GeneratedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20GeneratedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20GeneratedSession struct {
	Contract     *Erc20Generated   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20GeneratedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20GeneratedCallerSession struct {
	Contract *Erc20GeneratedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Erc20GeneratedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20GeneratedTransactorSession struct {
	Contract     *Erc20GeneratedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Erc20GeneratedRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20GeneratedRaw struct {
	Contract *Erc20Generated // Generic contract binding to access the raw methods on
}

// Erc20GeneratedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20GeneratedCallerRaw struct {
	Contract *Erc20GeneratedCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20GeneratedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20GeneratedTransactorRaw struct {
	Contract *Erc20GeneratedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Generated creates a new instance of Erc20Generated, bound to a specific deployed contract.
func NewErc20Generated(address common.Address, backend bind.ContractBackend) (*Erc20Generated, error) {
	contract, err := bindErc20Generated(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Generated{Erc20GeneratedCaller: Erc20GeneratedCaller{contract: contract}, Erc20GeneratedTransactor: Erc20GeneratedTransactor{contract: contract}, Erc20GeneratedFilterer: Erc20GeneratedFilterer{contract: contract}}, nil
}

// NewErc20GeneratedCaller creates a new read-only instance of Erc20Generated, bound to a specific deployed contract.
func NewErc20GeneratedCaller(address common.Address, caller bind.ContractCaller) (*Erc20GeneratedCaller, error) {
	contract, err := bindErc20Generated(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20GeneratedCaller{contract: contract}, nil
}

// NewErc20GeneratedTransactor creates a new write-only instance of Erc20Generated, bound to a specific deployed contract.
func NewErc20GeneratedTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20GeneratedTransactor, error) {
	contract, err := bindErc20Generated(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20GeneratedTransactor{contract: contract}, nil
}

// NewErc20GeneratedFilterer creates a new log filterer instance of Erc20Generated, bound to a specific deployed contract.
func NewErc20GeneratedFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20GeneratedFilterer, error) {
	contract, err := bindErc20Generated(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20GeneratedFilterer{contract: contract}, nil
}

// bindErc20Generated binds a generic wrapper to an already deployed contract.
func bindErc20Generated(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20GeneratedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Generated *Erc20GeneratedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Generated.Contract.Erc20GeneratedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Generated *Erc20GeneratedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Generated.Contract.Erc20GeneratedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Generated *Erc20GeneratedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Generated.Contract.Erc20GeneratedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Generated *Erc20GeneratedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Generated.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Generated *Erc20GeneratedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Generated.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Generated *Erc20GeneratedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Generated.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20Generated *Erc20GeneratedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Generated.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20Generated *Erc20GeneratedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20Generated.Contract.Allowance(&_Erc20Generated.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20Generated *Erc20GeneratedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20Generated.Contract.Allowance(&_Erc20Generated.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20Generated *Erc20GeneratedCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Generated.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20Generated *Erc20GeneratedSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20Generated.Contract.BalanceOf(&_Erc20Generated.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20Generated *Erc20GeneratedCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20Generated.Contract.BalanceOf(&_Erc20Generated.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20Generated *Erc20GeneratedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20Generated.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20Generated *Erc20GeneratedSession) Decimals() (uint8, error) {
	return _Erc20Generated.Contract.Decimals(&_Erc20Generated.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20Generated *Erc20GeneratedCallerSession) Decimals() (uint8, error) {
	return _Erc20Generated.Contract.Decimals(&_Erc20Generated.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Erc20Generated *Erc20GeneratedCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20Generated.contract.Call(opts, &out, "name")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Erc20Generated *Erc20GeneratedSession) Name() ([32]byte, error) {
	return _Erc20Generated.Contract.Name(&_Erc20Generated.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Erc20Generated *Erc20GeneratedCallerSession) Name() ([32]byte, error) {
	return _Erc20Generated.Contract.Name(&_Erc20Generated.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes32)
func (_Erc20Generated *Erc20GeneratedCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20Generated.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes32)
func (_Erc20Generated *Erc20GeneratedSession) Symbol() ([32]byte, error) {
	return _Erc20Generated.Contract.Symbol(&_Erc20Generated.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes32)
func (_Erc20Generated *Erc20GeneratedCallerSession) Symbol() ([32]byte, error) {
	return _Erc20Generated.Contract.Symbol(&_Erc20Generated.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20Generated *Erc20GeneratedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Generated.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20Generated *Erc20GeneratedSession) TotalSupply() (*big.Int, error) {
	return _Erc20Generated.Contract.TotalSupply(&_Erc20Generated.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20Generated *Erc20GeneratedCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20Generated.Contract.TotalSupply(&_Erc20Generated.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20Generated *Erc20GeneratedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Generated.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20Generated *Erc20GeneratedSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Generated.Contract.Approve(&_Erc20Generated.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20Generated *Erc20GeneratedTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Generated.Contract.Approve(&_Erc20Generated.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20Generated *Erc20GeneratedTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Generated.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20Generated *Erc20GeneratedSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Generated.Contract.Transfer(&_Erc20Generated.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20Generated *Erc20GeneratedTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Generated.Contract.Transfer(&_Erc20Generated.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20Generated *Erc20GeneratedTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Generated.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20Generated *Erc20GeneratedSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Generated.Contract.TransferFrom(&_Erc20Generated.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20Generated *Erc20GeneratedTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Generated.Contract.TransferFrom(&_Erc20Generated.TransactOpts, sender, recipient, amount)
}

// Erc20GeneratedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20Generated contract.
type Erc20GeneratedApprovalIterator struct {
	Event *Erc20GeneratedApproval // Event containing the contract specifics and raw log

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
func (it *Erc20GeneratedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20GeneratedApproval)
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
		it.Event = new(Erc20GeneratedApproval)
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
func (it *Erc20GeneratedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20GeneratedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20GeneratedApproval represents a Approval event raised by the Erc20Generated contract.
type Erc20GeneratedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Generated *Erc20GeneratedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20GeneratedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20Generated.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20GeneratedApprovalIterator{contract: _Erc20Generated.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Generated *Erc20GeneratedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20GeneratedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20Generated.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20GeneratedApproval)
				if err := _Erc20Generated.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Generated *Erc20GeneratedFilterer) ParseApproval(log types.Log) (*Erc20GeneratedApproval, error) {
	event := new(Erc20GeneratedApproval)
	if err := _Erc20Generated.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20GeneratedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20Generated contract.
type Erc20GeneratedTransferIterator struct {
	Event *Erc20GeneratedTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20GeneratedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20GeneratedTransfer)
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
		it.Event = new(Erc20GeneratedTransfer)
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
func (it *Erc20GeneratedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20GeneratedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20GeneratedTransfer represents a Transfer event raised by the Erc20Generated contract.
type Erc20GeneratedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Generated *Erc20GeneratedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20GeneratedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20Generated.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20GeneratedTransferIterator{contract: _Erc20Generated.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Generated *Erc20GeneratedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20GeneratedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20Generated.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20GeneratedTransfer)
				if err := _Erc20Generated.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Generated *Erc20GeneratedFilterer) ParseTransfer(log types.Log) (*Erc20GeneratedTransfer, error) {
	event := new(Erc20GeneratedTransfer)
	if err := _Erc20Generated.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
