// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bex

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

// CosmosCoin is an auto generated low-level Go binding around an user-defined struct.
type CosmosCoin struct {
	Amount *big.Int
	Denom  string
}

// RewardModuleMetaData contains all meta data concerning the RewardModule contract.
var RewardModuleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getCurrentRewards\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositorWithdrawAddress\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutstandingRewards\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDepositorWithdrawAddress\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllDepositorRewards\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawDepositorRewards\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawDepositorRewardsTo\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"InitializeDeposit\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"shares\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCosmos.Coin\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDepositorWithdrawAddress\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawDepositRewards\",\"inputs\":[{\"name\":\"rewardReceiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"rewardAmount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false}]",
}

// RewardModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardModuleMetaData.ABI instead.
var RewardModuleABI = RewardModuleMetaData.ABI

// RewardModule is an auto generated Go binding around an Ethereum contract.
type RewardModule struct {
	RewardModuleCaller     // Read-only binding to the contract
	RewardModuleTransactor // Write-only binding to the contract
	RewardModuleFilterer   // Log filterer for contract events
}

// RewardModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardModuleSession struct {
	Contract     *RewardModule     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardModuleCallerSession struct {
	Contract *RewardModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RewardModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardModuleTransactorSession struct {
	Contract     *RewardModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RewardModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardModuleRaw struct {
	Contract *RewardModule // Generic contract binding to access the raw methods on
}

// RewardModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardModuleCallerRaw struct {
	Contract *RewardModuleCaller // Generic read-only contract binding to access the raw methods on
}

// RewardModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardModuleTransactorRaw struct {
	Contract *RewardModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardModule creates a new instance of RewardModule, bound to a specific deployed contract.
func NewRewardModule(address common.Address, backend bind.ContractBackend) (*RewardModule, error) {
	contract, err := bindRewardModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardModule{RewardModuleCaller: RewardModuleCaller{contract: contract}, RewardModuleTransactor: RewardModuleTransactor{contract: contract}, RewardModuleFilterer: RewardModuleFilterer{contract: contract}}, nil
}

// NewRewardModuleCaller creates a new read-only instance of RewardModule, bound to a specific deployed contract.
func NewRewardModuleCaller(address common.Address, caller bind.ContractCaller) (*RewardModuleCaller, error) {
	contract, err := bindRewardModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardModuleCaller{contract: contract}, nil
}

// NewRewardModuleTransactor creates a new write-only instance of RewardModule, bound to a specific deployed contract.
func NewRewardModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardModuleTransactor, error) {
	contract, err := bindRewardModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardModuleTransactor{contract: contract}, nil
}

// NewRewardModuleFilterer creates a new log filterer instance of RewardModule, bound to a specific deployed contract.
func NewRewardModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardModuleFilterer, error) {
	contract, err := bindRewardModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardModuleFilterer{contract: contract}, nil
}

// bindRewardModule binds a generic wrapper to an already deployed contract.
func bindRewardModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardModule *RewardModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardModule.Contract.RewardModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardModule *RewardModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardModule.Contract.RewardModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardModule *RewardModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardModule.Contract.RewardModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardModule *RewardModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardModule *RewardModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardModule *RewardModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardModule.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentRewards is a free data retrieval call binding the contract method 0x8cb3ca0b.
//
// Solidity: function getCurrentRewards(address depositor, address receiver) view returns((uint256,string)[])
func (_RewardModule *RewardModuleCaller) GetCurrentRewards(opts *bind.CallOpts, depositor common.Address, receiver common.Address) ([]CosmosCoin, error) {
	var out []interface{}
	err := _RewardModule.contract.Call(opts, &out, "getCurrentRewards", depositor, receiver)

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetCurrentRewards is a free data retrieval call binding the contract method 0x8cb3ca0b.
//
// Solidity: function getCurrentRewards(address depositor, address receiver) view returns((uint256,string)[])
func (_RewardModule *RewardModuleSession) GetCurrentRewards(depositor common.Address, receiver common.Address) ([]CosmosCoin, error) {
	return _RewardModule.Contract.GetCurrentRewards(&_RewardModule.CallOpts, depositor, receiver)
}

// GetCurrentRewards is a free data retrieval call binding the contract method 0x8cb3ca0b.
//
// Solidity: function getCurrentRewards(address depositor, address receiver) view returns((uint256,string)[])
func (_RewardModule *RewardModuleCallerSession) GetCurrentRewards(depositor common.Address, receiver common.Address) ([]CosmosCoin, error) {
	return _RewardModule.Contract.GetCurrentRewards(&_RewardModule.CallOpts, depositor, receiver)
}

// GetDepositorWithdrawAddress is a free data retrieval call binding the contract method 0x54abed38.
//
// Solidity: function getDepositorWithdrawAddress(address depositor) view returns(address)
func (_RewardModule *RewardModuleCaller) GetDepositorWithdrawAddress(opts *bind.CallOpts, depositor common.Address) (common.Address, error) {
	var out []interface{}
	err := _RewardModule.contract.Call(opts, &out, "getDepositorWithdrawAddress", depositor)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDepositorWithdrawAddress is a free data retrieval call binding the contract method 0x54abed38.
//
// Solidity: function getDepositorWithdrawAddress(address depositor) view returns(address)
func (_RewardModule *RewardModuleSession) GetDepositorWithdrawAddress(depositor common.Address) (common.Address, error) {
	return _RewardModule.Contract.GetDepositorWithdrawAddress(&_RewardModule.CallOpts, depositor)
}

// GetDepositorWithdrawAddress is a free data retrieval call binding the contract method 0x54abed38.
//
// Solidity: function getDepositorWithdrawAddress(address depositor) view returns(address)
func (_RewardModule *RewardModuleCallerSession) GetDepositorWithdrawAddress(depositor common.Address) (common.Address, error) {
	return _RewardModule.Contract.GetDepositorWithdrawAddress(&_RewardModule.CallOpts, depositor)
}

// GetOutstandingRewards is a free data retrieval call binding the contract method 0xbce3b836.
//
// Solidity: function getOutstandingRewards(address receiver) view returns((uint256,string)[])
func (_RewardModule *RewardModuleCaller) GetOutstandingRewards(opts *bind.CallOpts, receiver common.Address) ([]CosmosCoin, error) {
	var out []interface{}
	err := _RewardModule.contract.Call(opts, &out, "getOutstandingRewards", receiver)

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetOutstandingRewards is a free data retrieval call binding the contract method 0xbce3b836.
//
// Solidity: function getOutstandingRewards(address receiver) view returns((uint256,string)[])
func (_RewardModule *RewardModuleSession) GetOutstandingRewards(receiver common.Address) ([]CosmosCoin, error) {
	return _RewardModule.Contract.GetOutstandingRewards(&_RewardModule.CallOpts, receiver)
}

// GetOutstandingRewards is a free data retrieval call binding the contract method 0xbce3b836.
//
// Solidity: function getOutstandingRewards(address receiver) view returns((uint256,string)[])
func (_RewardModule *RewardModuleCallerSession) GetOutstandingRewards(receiver common.Address) ([]CosmosCoin, error) {
	return _RewardModule.Contract.GetOutstandingRewards(&_RewardModule.CallOpts, receiver)
}

// SetDepositorWithdrawAddress is a paid mutator transaction binding the contract method 0x56c4d0db.
//
// Solidity: function setDepositorWithdrawAddress(address withdrawAddress) returns(bool)
func (_RewardModule *RewardModuleTransactor) SetDepositorWithdrawAddress(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _RewardModule.contract.Transact(opts, "setDepositorWithdrawAddress", withdrawAddress)
}

// SetDepositorWithdrawAddress is a paid mutator transaction binding the contract method 0x56c4d0db.
//
// Solidity: function setDepositorWithdrawAddress(address withdrawAddress) returns(bool)
func (_RewardModule *RewardModuleSession) SetDepositorWithdrawAddress(withdrawAddress common.Address) (*types.Transaction, error) {
	return _RewardModule.Contract.SetDepositorWithdrawAddress(&_RewardModule.TransactOpts, withdrawAddress)
}

// SetDepositorWithdrawAddress is a paid mutator transaction binding the contract method 0x56c4d0db.
//
// Solidity: function setDepositorWithdrawAddress(address withdrawAddress) returns(bool)
func (_RewardModule *RewardModuleTransactorSession) SetDepositorWithdrawAddress(withdrawAddress common.Address) (*types.Transaction, error) {
	return _RewardModule.Contract.SetDepositorWithdrawAddress(&_RewardModule.TransactOpts, withdrawAddress)
}

// WithdrawAllDepositorRewards is a paid mutator transaction binding the contract method 0xc02e8929.
//
// Solidity: function withdrawAllDepositorRewards(address receiver) returns((uint256,string)[])
func (_RewardModule *RewardModuleTransactor) WithdrawAllDepositorRewards(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _RewardModule.contract.Transact(opts, "withdrawAllDepositorRewards", receiver)
}

// WithdrawAllDepositorRewards is a paid mutator transaction binding the contract method 0xc02e8929.
//
// Solidity: function withdrawAllDepositorRewards(address receiver) returns((uint256,string)[])
func (_RewardModule *RewardModuleSession) WithdrawAllDepositorRewards(receiver common.Address) (*types.Transaction, error) {
	return _RewardModule.Contract.WithdrawAllDepositorRewards(&_RewardModule.TransactOpts, receiver)
}

// WithdrawAllDepositorRewards is a paid mutator transaction binding the contract method 0xc02e8929.
//
// Solidity: function withdrawAllDepositorRewards(address receiver) returns((uint256,string)[])
func (_RewardModule *RewardModuleTransactorSession) WithdrawAllDepositorRewards(receiver common.Address) (*types.Transaction, error) {
	return _RewardModule.Contract.WithdrawAllDepositorRewards(&_RewardModule.TransactOpts, receiver)
}

// WithdrawDepositorRewards is a paid mutator transaction binding the contract method 0x0d2c9ec8.
//
// Solidity: function withdrawDepositorRewards(address receiver, uint256 amount) returns((uint256,string)[])
func (_RewardModule *RewardModuleTransactor) WithdrawDepositorRewards(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardModule.contract.Transact(opts, "withdrawDepositorRewards", receiver, amount)
}

// WithdrawDepositorRewards is a paid mutator transaction binding the contract method 0x0d2c9ec8.
//
// Solidity: function withdrawDepositorRewards(address receiver, uint256 amount) returns((uint256,string)[])
func (_RewardModule *RewardModuleSession) WithdrawDepositorRewards(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardModule.Contract.WithdrawDepositorRewards(&_RewardModule.TransactOpts, receiver, amount)
}

// WithdrawDepositorRewards is a paid mutator transaction binding the contract method 0x0d2c9ec8.
//
// Solidity: function withdrawDepositorRewards(address receiver, uint256 amount) returns((uint256,string)[])
func (_RewardModule *RewardModuleTransactorSession) WithdrawDepositorRewards(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardModule.Contract.WithdrawDepositorRewards(&_RewardModule.TransactOpts, receiver, amount)
}

// WithdrawDepositorRewardsTo is a paid mutator transaction binding the contract method 0x3771f642.
//
// Solidity: function withdrawDepositorRewardsTo(address receiver, address recipient, uint256 amount) returns((uint256,string)[])
func (_RewardModule *RewardModuleTransactor) WithdrawDepositorRewardsTo(opts *bind.TransactOpts, receiver common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardModule.contract.Transact(opts, "withdrawDepositorRewardsTo", receiver, recipient, amount)
}

// WithdrawDepositorRewardsTo is a paid mutator transaction binding the contract method 0x3771f642.
//
// Solidity: function withdrawDepositorRewardsTo(address receiver, address recipient, uint256 amount) returns((uint256,string)[])
func (_RewardModule *RewardModuleSession) WithdrawDepositorRewardsTo(receiver common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardModule.Contract.WithdrawDepositorRewardsTo(&_RewardModule.TransactOpts, receiver, recipient, amount)
}

// WithdrawDepositorRewardsTo is a paid mutator transaction binding the contract method 0x3771f642.
//
// Solidity: function withdrawDepositorRewardsTo(address receiver, address recipient, uint256 amount) returns((uint256,string)[])
func (_RewardModule *RewardModuleTransactorSession) WithdrawDepositorRewardsTo(receiver common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardModule.Contract.WithdrawDepositorRewardsTo(&_RewardModule.TransactOpts, receiver, recipient, amount)
}

// RewardModuleInitializeDepositIterator is returned from FilterInitializeDeposit and is used to iterate over the raw logs and unpacked data for InitializeDeposit events raised by the RewardModule contract.
type RewardModuleInitializeDepositIterator struct {
	Event *RewardModuleInitializeDeposit // Event containing the contract specifics and raw log

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
func (it *RewardModuleInitializeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardModuleInitializeDeposit)
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
		it.Event = new(RewardModuleInitializeDeposit)
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
func (it *RewardModuleInitializeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardModuleInitializeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardModuleInitializeDeposit represents a InitializeDeposit event raised by the RewardModule contract.
type RewardModuleInitializeDeposit struct {
	Caller    common.Address
	Depositor common.Address
	Assets    []CosmosCoin
	Shares    CosmosCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitializeDeposit is a free log retrieval operation binding the contract event 0xd924fa351aaddb8d643abd0e9649dba91fe2ea3597e3d50ea5b35c5b590504cc.
//
// Solidity: event InitializeDeposit(address indexed caller, address indexed depositor, (uint256,string)[] assets, (uint256,string) shares)
func (_RewardModule *RewardModuleFilterer) FilterInitializeDeposit(opts *bind.FilterOpts, caller []common.Address, depositor []common.Address) (*RewardModuleInitializeDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _RewardModule.contract.FilterLogs(opts, "InitializeDeposit", callerRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &RewardModuleInitializeDepositIterator{contract: _RewardModule.contract, event: "InitializeDeposit", logs: logs, sub: sub}, nil
}

// WatchInitializeDeposit is a free log subscription operation binding the contract event 0xd924fa351aaddb8d643abd0e9649dba91fe2ea3597e3d50ea5b35c5b590504cc.
//
// Solidity: event InitializeDeposit(address indexed caller, address indexed depositor, (uint256,string)[] assets, (uint256,string) shares)
func (_RewardModule *RewardModuleFilterer) WatchInitializeDeposit(opts *bind.WatchOpts, sink chan<- *RewardModuleInitializeDeposit, caller []common.Address, depositor []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _RewardModule.contract.WatchLogs(opts, "InitializeDeposit", callerRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardModuleInitializeDeposit)
				if err := _RewardModule.contract.UnpackLog(event, "InitializeDeposit", log); err != nil {
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

// ParseInitializeDeposit is a log parse operation binding the contract event 0xd924fa351aaddb8d643abd0e9649dba91fe2ea3597e3d50ea5b35c5b590504cc.
//
// Solidity: event InitializeDeposit(address indexed caller, address indexed depositor, (uint256,string)[] assets, (uint256,string) shares)
func (_RewardModule *RewardModuleFilterer) ParseInitializeDeposit(log types.Log) (*RewardModuleInitializeDeposit, error) {
	event := new(RewardModuleInitializeDeposit)
	if err := _RewardModule.contract.UnpackLog(event, "InitializeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardModuleSetDepositorWithdrawAddressIterator is returned from FilterSetDepositorWithdrawAddress and is used to iterate over the raw logs and unpacked data for SetDepositorWithdrawAddress events raised by the RewardModule contract.
type RewardModuleSetDepositorWithdrawAddressIterator struct {
	Event *RewardModuleSetDepositorWithdrawAddress // Event containing the contract specifics and raw log

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
func (it *RewardModuleSetDepositorWithdrawAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardModuleSetDepositorWithdrawAddress)
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
		it.Event = new(RewardModuleSetDepositorWithdrawAddress)
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
func (it *RewardModuleSetDepositorWithdrawAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardModuleSetDepositorWithdrawAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardModuleSetDepositorWithdrawAddress represents a SetDepositorWithdrawAddress event raised by the RewardModule contract.
type RewardModuleSetDepositorWithdrawAddress struct {
	Depositor       common.Address
	WithdrawAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetDepositorWithdrawAddress is a free log retrieval operation binding the contract event 0x0de02018d2d8e05a493bcec83d64d09b0bbe4320855afc0b410c0af84f3b6241.
//
// Solidity: event SetDepositorWithdrawAddress(address indexed depositor, address indexed withdrawAddress)
func (_RewardModule *RewardModuleFilterer) FilterSetDepositorWithdrawAddress(opts *bind.FilterOpts, depositor []common.Address, withdrawAddress []common.Address) (*RewardModuleSetDepositorWithdrawAddressIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var withdrawAddressRule []interface{}
	for _, withdrawAddressItem := range withdrawAddress {
		withdrawAddressRule = append(withdrawAddressRule, withdrawAddressItem)
	}

	logs, sub, err := _RewardModule.contract.FilterLogs(opts, "SetDepositorWithdrawAddress", depositorRule, withdrawAddressRule)
	if err != nil {
		return nil, err
	}
	return &RewardModuleSetDepositorWithdrawAddressIterator{contract: _RewardModule.contract, event: "SetDepositorWithdrawAddress", logs: logs, sub: sub}, nil
}

// WatchSetDepositorWithdrawAddress is a free log subscription operation binding the contract event 0x0de02018d2d8e05a493bcec83d64d09b0bbe4320855afc0b410c0af84f3b6241.
//
// Solidity: event SetDepositorWithdrawAddress(address indexed depositor, address indexed withdrawAddress)
func (_RewardModule *RewardModuleFilterer) WatchSetDepositorWithdrawAddress(opts *bind.WatchOpts, sink chan<- *RewardModuleSetDepositorWithdrawAddress, depositor []common.Address, withdrawAddress []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var withdrawAddressRule []interface{}
	for _, withdrawAddressItem := range withdrawAddress {
		withdrawAddressRule = append(withdrawAddressRule, withdrawAddressItem)
	}

	logs, sub, err := _RewardModule.contract.WatchLogs(opts, "SetDepositorWithdrawAddress", depositorRule, withdrawAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardModuleSetDepositorWithdrawAddress)
				if err := _RewardModule.contract.UnpackLog(event, "SetDepositorWithdrawAddress", log); err != nil {
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

// ParseSetDepositorWithdrawAddress is a log parse operation binding the contract event 0x0de02018d2d8e05a493bcec83d64d09b0bbe4320855afc0b410c0af84f3b6241.
//
// Solidity: event SetDepositorWithdrawAddress(address indexed depositor, address indexed withdrawAddress)
func (_RewardModule *RewardModuleFilterer) ParseSetDepositorWithdrawAddress(log types.Log) (*RewardModuleSetDepositorWithdrawAddress, error) {
	event := new(RewardModuleSetDepositorWithdrawAddress)
	if err := _RewardModule.contract.UnpackLog(event, "SetDepositorWithdrawAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardModuleWithdrawDepositRewardsIterator is returned from FilterWithdrawDepositRewards and is used to iterate over the raw logs and unpacked data for WithdrawDepositRewards events raised by the RewardModule contract.
type RewardModuleWithdrawDepositRewardsIterator struct {
	Event *RewardModuleWithdrawDepositRewards // Event containing the contract specifics and raw log

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
func (it *RewardModuleWithdrawDepositRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardModuleWithdrawDepositRewards)
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
		it.Event = new(RewardModuleWithdrawDepositRewards)
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
func (it *RewardModuleWithdrawDepositRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardModuleWithdrawDepositRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardModuleWithdrawDepositRewards represents a WithdrawDepositRewards event raised by the RewardModule contract.
type RewardModuleWithdrawDepositRewards struct {
	RewardReceiver  common.Address
	Withdrawer      common.Address
	RewardRecipient common.Address
	RewardAmount    []CosmosCoin
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawDepositRewards is a free log retrieval operation binding the contract event 0xeaf9ad68989c4779cdb8274ed825acd8dd664e333e1168e03e7fb4766bc514a4.
//
// Solidity: event WithdrawDepositRewards(address indexed rewardReceiver, address indexed withdrawer, address indexed rewardRecipient, (uint256,string)[] rewardAmount)
func (_RewardModule *RewardModuleFilterer) FilterWithdrawDepositRewards(opts *bind.FilterOpts, rewardReceiver []common.Address, withdrawer []common.Address, rewardRecipient []common.Address) (*RewardModuleWithdrawDepositRewardsIterator, error) {

	var rewardReceiverRule []interface{}
	for _, rewardReceiverItem := range rewardReceiver {
		rewardReceiverRule = append(rewardReceiverRule, rewardReceiverItem)
	}
	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}
	var rewardRecipientRule []interface{}
	for _, rewardRecipientItem := range rewardRecipient {
		rewardRecipientRule = append(rewardRecipientRule, rewardRecipientItem)
	}

	logs, sub, err := _RewardModule.contract.FilterLogs(opts, "WithdrawDepositRewards", rewardReceiverRule, withdrawerRule, rewardRecipientRule)
	if err != nil {
		return nil, err
	}
	return &RewardModuleWithdrawDepositRewardsIterator{contract: _RewardModule.contract, event: "WithdrawDepositRewards", logs: logs, sub: sub}, nil
}

// WatchWithdrawDepositRewards is a free log subscription operation binding the contract event 0xeaf9ad68989c4779cdb8274ed825acd8dd664e333e1168e03e7fb4766bc514a4.
//
// Solidity: event WithdrawDepositRewards(address indexed rewardReceiver, address indexed withdrawer, address indexed rewardRecipient, (uint256,string)[] rewardAmount)
func (_RewardModule *RewardModuleFilterer) WatchWithdrawDepositRewards(opts *bind.WatchOpts, sink chan<- *RewardModuleWithdrawDepositRewards, rewardReceiver []common.Address, withdrawer []common.Address, rewardRecipient []common.Address) (event.Subscription, error) {

	var rewardReceiverRule []interface{}
	for _, rewardReceiverItem := range rewardReceiver {
		rewardReceiverRule = append(rewardReceiverRule, rewardReceiverItem)
	}
	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}
	var rewardRecipientRule []interface{}
	for _, rewardRecipientItem := range rewardRecipient {
		rewardRecipientRule = append(rewardRecipientRule, rewardRecipientItem)
	}

	logs, sub, err := _RewardModule.contract.WatchLogs(opts, "WithdrawDepositRewards", rewardReceiverRule, withdrawerRule, rewardRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardModuleWithdrawDepositRewards)
				if err := _RewardModule.contract.UnpackLog(event, "WithdrawDepositRewards", log); err != nil {
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

// ParseWithdrawDepositRewards is a log parse operation binding the contract event 0xeaf9ad68989c4779cdb8274ed825acd8dd664e333e1168e03e7fb4766bc514a4.
//
// Solidity: event WithdrawDepositRewards(address indexed rewardReceiver, address indexed withdrawer, address indexed rewardRecipient, (uint256,string)[] rewardAmount)
func (_RewardModule *RewardModuleFilterer) ParseWithdrawDepositRewards(log types.Log) (*RewardModuleWithdrawDepositRewards, error) {
	event := new(RewardModuleWithdrawDepositRewards)
	if err := _RewardModule.contract.UnpackLog(event, "WithdrawDepositRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
