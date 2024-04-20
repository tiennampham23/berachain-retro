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

// IERC20DexModuleAssetWeight is an auto generated low-level Go binding around an user-defined struct.
type IERC20DexModuleAssetWeight struct {
	Asset  common.Address
	Weight *big.Int
}

// IERC20DexModuleBatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
type IERC20DexModuleBatchSwapStep struct {
	PoolId    common.Address
	AssetIn   common.Address
	AmountIn  *big.Int
	AssetOut  common.Address
	AmountOut *big.Int
	UserData  []byte
}

// IERC20DexModulePoolOptions is an auto generated low-level Go binding around an user-defined struct.
type IERC20DexModulePoolOptions struct {
	Weights []IERC20DexModuleAssetWeight
	SwapFee *big.Int
}

// BexGeneratedMetaData contains all meta data concerning the BexGenerated contract.
var BexGeneratedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assetsIn\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"shares\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shareAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"liquidity\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIERC20DexModule.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"poolId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIERC20DexModule.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"batchSwap\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"assetsIn\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"poolType\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIERC20DexModule.AssetWeight[]\",\"name\":\"weights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"internalType\":\"structIERC20DexModule.PoolOptions\",\"name\":\"options\",\"type\":\"tuple\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"getExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getLiquidity\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"asset\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolOptions\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIERC20DexModule.AssetWeight[]\",\"name\":\"weights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"internalType\":\"structIERC20DexModule.PoolOptions\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"getPreviewAddLiquidityNoSwap\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"shares\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shareAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"liqOut\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"liquidity\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"getPreviewAddLiquidityStaticPrice\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"shares\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shareAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"liqOut\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIERC20DexModule.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"poolId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIERC20DexModule.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"}],\"name\":\"getPreviewBatchSwap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPreviewBurnShares\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"getPreviewSharesForLiquidity\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"shares\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shareAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"liquidity\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPreviewSharesForSingleSidedLiquidityRequest\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIERC20DexModule.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"getPreviewSwapExact\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"}],\"name\":\"getRemoveLiquidityExactAmountOut\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sharesIn\",\"type\":\"uint256\"}],\"name\":\"getRemoveLiquidityOneSideOut\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getTotalShares\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityBurningShares\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"liquidity\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sharesIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSharesIn\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityExactAmount\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"shares\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shareAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"liquidity\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIERC20DexModule.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"poolId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BexGeneratedABI is the input ABI used to generate the binding from.
// Deprecated: Use BexGeneratedMetaData.ABI instead.
var BexGeneratedABI = BexGeneratedMetaData.ABI

// BexGenerated is an auto generated Go binding around an Ethereum contract.
type BexGenerated struct {
	BexGeneratedCaller     // Read-only binding to the contract
	BexGeneratedTransactor // Write-only binding to the contract
	BexGeneratedFilterer   // Log filterer for contract events
}

// BexGeneratedCaller is an auto generated read-only Go binding around an Ethereum contract.
type BexGeneratedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BexGeneratedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BexGeneratedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BexGeneratedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BexGeneratedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BexGeneratedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BexGeneratedSession struct {
	Contract     *BexGenerated     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BexGeneratedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BexGeneratedCallerSession struct {
	Contract *BexGeneratedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BexGeneratedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BexGeneratedTransactorSession struct {
	Contract     *BexGeneratedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BexGeneratedRaw is an auto generated low-level Go binding around an Ethereum contract.
type BexGeneratedRaw struct {
	Contract *BexGenerated // Generic contract binding to access the raw methods on
}

// BexGeneratedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BexGeneratedCallerRaw struct {
	Contract *BexGeneratedCaller // Generic read-only contract binding to access the raw methods on
}

// BexGeneratedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BexGeneratedTransactorRaw struct {
	Contract *BexGeneratedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBexGenerated creates a new instance of BexGenerated, bound to a specific deployed contract.
func NewBexGenerated(address common.Address, backend bind.ContractBackend) (*BexGenerated, error) {
	contract, err := bindBexGenerated(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BexGenerated{BexGeneratedCaller: BexGeneratedCaller{contract: contract}, BexGeneratedTransactor: BexGeneratedTransactor{contract: contract}, BexGeneratedFilterer: BexGeneratedFilterer{contract: contract}}, nil
}

// NewBexGeneratedCaller creates a new read-only instance of BexGenerated, bound to a specific deployed contract.
func NewBexGeneratedCaller(address common.Address, caller bind.ContractCaller) (*BexGeneratedCaller, error) {
	contract, err := bindBexGenerated(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BexGeneratedCaller{contract: contract}, nil
}

// NewBexGeneratedTransactor creates a new write-only instance of BexGenerated, bound to a specific deployed contract.
func NewBexGeneratedTransactor(address common.Address, transactor bind.ContractTransactor) (*BexGeneratedTransactor, error) {
	contract, err := bindBexGenerated(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BexGeneratedTransactor{contract: contract}, nil
}

// NewBexGeneratedFilterer creates a new log filterer instance of BexGenerated, bound to a specific deployed contract.
func NewBexGeneratedFilterer(address common.Address, filterer bind.ContractFilterer) (*BexGeneratedFilterer, error) {
	contract, err := bindBexGenerated(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BexGeneratedFilterer{contract: contract}, nil
}

// bindBexGenerated binds a generic wrapper to an already deployed contract.
func bindBexGenerated(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BexGeneratedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BexGenerated *BexGeneratedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BexGenerated.Contract.BexGeneratedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BexGenerated *BexGeneratedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BexGenerated.Contract.BexGeneratedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BexGenerated *BexGeneratedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BexGenerated.Contract.BexGeneratedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BexGenerated *BexGeneratedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BexGenerated.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BexGenerated *BexGeneratedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BexGenerated.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BexGenerated *BexGeneratedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BexGenerated.Contract.contract.Transact(opts, method, params...)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0x2575c021.
//
// Solidity: function getExchangeRate(address pool, address baseAsset, address quoteAsset) view returns(uint256)
func (_BexGenerated *BexGeneratedCaller) GetExchangeRate(opts *bind.CallOpts, pool common.Address, baseAsset common.Address, quoteAsset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getExchangeRate", pool, baseAsset, quoteAsset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0x2575c021.
//
// Solidity: function getExchangeRate(address pool, address baseAsset, address quoteAsset) view returns(uint256)
func (_BexGenerated *BexGeneratedSession) GetExchangeRate(pool common.Address, baseAsset common.Address, quoteAsset common.Address) (*big.Int, error) {
	return _BexGenerated.Contract.GetExchangeRate(&_BexGenerated.CallOpts, pool, baseAsset, quoteAsset)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0x2575c021.
//
// Solidity: function getExchangeRate(address pool, address baseAsset, address quoteAsset) view returns(uint256)
func (_BexGenerated *BexGeneratedCallerSession) GetExchangeRate(pool common.Address, baseAsset common.Address, quoteAsset common.Address) (*big.Int, error) {
	return _BexGenerated.Contract.GetExchangeRate(&_BexGenerated.CallOpts, pool, baseAsset, quoteAsset)
}

// GetLiquidity is a free data retrieval call binding the contract method 0xa747b93b.
//
// Solidity: function getLiquidity(address pool) view returns(address[] asset, uint256[] amounts)
func (_BexGenerated *BexGeneratedCaller) GetLiquidity(opts *bind.CallOpts, pool common.Address) (struct {
	Asset   []common.Address
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getLiquidity", pool)

	outstruct := new(struct {
		Asset   []common.Address
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Asset = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetLiquidity is a free data retrieval call binding the contract method 0xa747b93b.
//
// Solidity: function getLiquidity(address pool) view returns(address[] asset, uint256[] amounts)
func (_BexGenerated *BexGeneratedSession) GetLiquidity(pool common.Address) (struct {
	Asset   []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetLiquidity(&_BexGenerated.CallOpts, pool)
}

// GetLiquidity is a free data retrieval call binding the contract method 0xa747b93b.
//
// Solidity: function getLiquidity(address pool) view returns(address[] asset, uint256[] amounts)
func (_BexGenerated *BexGeneratedCallerSession) GetLiquidity(pool common.Address) (struct {
	Asset   []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetLiquidity(&_BexGenerated.CallOpts, pool)
}

// GetPoolName is a free data retrieval call binding the contract method 0x72afbc30.
//
// Solidity: function getPoolName(address pool) view returns(string)
func (_BexGenerated *BexGeneratedCaller) GetPoolName(opts *bind.CallOpts, pool common.Address) (string, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getPoolName", pool)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName is a free data retrieval call binding the contract method 0x72afbc30.
//
// Solidity: function getPoolName(address pool) view returns(string)
func (_BexGenerated *BexGeneratedSession) GetPoolName(pool common.Address) (string, error) {
	return _BexGenerated.Contract.GetPoolName(&_BexGenerated.CallOpts, pool)
}

// GetPoolName is a free data retrieval call binding the contract method 0x72afbc30.
//
// Solidity: function getPoolName(address pool) view returns(string)
func (_BexGenerated *BexGeneratedCallerSession) GetPoolName(pool common.Address) (string, error) {
	return _BexGenerated.Contract.GetPoolName(&_BexGenerated.CallOpts, pool)
}

// GetPoolOptions is a free data retrieval call binding the contract method 0x7ee84de4.
//
// Solidity: function getPoolOptions(address pool) view returns(((address,uint256)[],uint256))
func (_BexGenerated *BexGeneratedCaller) GetPoolOptions(opts *bind.CallOpts, pool common.Address) (IERC20DexModulePoolOptions, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getPoolOptions", pool)

	if err != nil {
		return *new(IERC20DexModulePoolOptions), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC20DexModulePoolOptions)).(*IERC20DexModulePoolOptions)

	return out0, err

}

// GetPoolOptions is a free data retrieval call binding the contract method 0x7ee84de4.
//
// Solidity: function getPoolOptions(address pool) view returns(((address,uint256)[],uint256))
func (_BexGenerated *BexGeneratedSession) GetPoolOptions(pool common.Address) (IERC20DexModulePoolOptions, error) {
	return _BexGenerated.Contract.GetPoolOptions(&_BexGenerated.CallOpts, pool)
}

// GetPoolOptions is a free data retrieval call binding the contract method 0x7ee84de4.
//
// Solidity: function getPoolOptions(address pool) view returns(((address,uint256)[],uint256))
func (_BexGenerated *BexGeneratedCallerSession) GetPoolOptions(pool common.Address) (IERC20DexModulePoolOptions, error) {
	return _BexGenerated.Contract.GetPoolOptions(&_BexGenerated.CallOpts, pool)
}

// GetPreviewAddLiquidityNoSwap is a free data retrieval call binding the contract method 0x89959aef.
//
// Solidity: function getPreviewAddLiquidityNoSwap(address pool, address[] assets, uint256[] amounts) view returns(address[] shares, uint256[] shareAmounts, address[] liqOut, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedCaller) GetPreviewAddLiquidityNoSwap(opts *bind.CallOpts, pool common.Address, assets []common.Address, amounts []*big.Int) (struct {
	Shares           []common.Address
	ShareAmounts     []*big.Int
	LiqOut           []common.Address
	LiquidityAmounts []*big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getPreviewAddLiquidityNoSwap", pool, assets, amounts)

	outstruct := new(struct {
		Shares           []common.Address
		ShareAmounts     []*big.Int
		LiqOut           []common.Address
		LiquidityAmounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.ShareAmounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LiqOut = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	outstruct.LiquidityAmounts = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPreviewAddLiquidityNoSwap is a free data retrieval call binding the contract method 0x89959aef.
//
// Solidity: function getPreviewAddLiquidityNoSwap(address pool, address[] assets, uint256[] amounts) view returns(address[] shares, uint256[] shareAmounts, address[] liqOut, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedSession) GetPreviewAddLiquidityNoSwap(pool common.Address, assets []common.Address, amounts []*big.Int) (struct {
	Shares           []common.Address
	ShareAmounts     []*big.Int
	LiqOut           []common.Address
	LiquidityAmounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewAddLiquidityNoSwap(&_BexGenerated.CallOpts, pool, assets, amounts)
}

// GetPreviewAddLiquidityNoSwap is a free data retrieval call binding the contract method 0x89959aef.
//
// Solidity: function getPreviewAddLiquidityNoSwap(address pool, address[] assets, uint256[] amounts) view returns(address[] shares, uint256[] shareAmounts, address[] liqOut, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedCallerSession) GetPreviewAddLiquidityNoSwap(pool common.Address, assets []common.Address, amounts []*big.Int) (struct {
	Shares           []common.Address
	ShareAmounts     []*big.Int
	LiqOut           []common.Address
	LiquidityAmounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewAddLiquidityNoSwap(&_BexGenerated.CallOpts, pool, assets, amounts)
}

// GetPreviewAddLiquidityStaticPrice is a free data retrieval call binding the contract method 0x38a2cc37.
//
// Solidity: function getPreviewAddLiquidityStaticPrice(address pool, address[] liquidity, uint256[] amounts) view returns(address[] shares, uint256[] shareAmounts, address[] liqOut, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedCaller) GetPreviewAddLiquidityStaticPrice(opts *bind.CallOpts, pool common.Address, liquidity []common.Address, amounts []*big.Int) (struct {
	Shares           []common.Address
	ShareAmounts     []*big.Int
	LiqOut           []common.Address
	LiquidityAmounts []*big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getPreviewAddLiquidityStaticPrice", pool, liquidity, amounts)

	outstruct := new(struct {
		Shares           []common.Address
		ShareAmounts     []*big.Int
		LiqOut           []common.Address
		LiquidityAmounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.ShareAmounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LiqOut = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	outstruct.LiquidityAmounts = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPreviewAddLiquidityStaticPrice is a free data retrieval call binding the contract method 0x38a2cc37.
//
// Solidity: function getPreviewAddLiquidityStaticPrice(address pool, address[] liquidity, uint256[] amounts) view returns(address[] shares, uint256[] shareAmounts, address[] liqOut, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedSession) GetPreviewAddLiquidityStaticPrice(pool common.Address, liquidity []common.Address, amounts []*big.Int) (struct {
	Shares           []common.Address
	ShareAmounts     []*big.Int
	LiqOut           []common.Address
	LiquidityAmounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewAddLiquidityStaticPrice(&_BexGenerated.CallOpts, pool, liquidity, amounts)
}

// GetPreviewAddLiquidityStaticPrice is a free data retrieval call binding the contract method 0x38a2cc37.
//
// Solidity: function getPreviewAddLiquidityStaticPrice(address pool, address[] liquidity, uint256[] amounts) view returns(address[] shares, uint256[] shareAmounts, address[] liqOut, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedCallerSession) GetPreviewAddLiquidityStaticPrice(pool common.Address, liquidity []common.Address, amounts []*big.Int) (struct {
	Shares           []common.Address
	ShareAmounts     []*big.Int
	LiqOut           []common.Address
	LiquidityAmounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewAddLiquidityStaticPrice(&_BexGenerated.CallOpts, pool, liquidity, amounts)
}

// GetPreviewBatchSwap is a free data retrieval call binding the contract method 0xe6ee1e0b.
//
// Solidity: function getPreviewBatchSwap(uint8 kind, (address,address,uint256,address,uint256,bytes)[] swaps) view returns(address asset, uint256 amount)
func (_BexGenerated *BexGeneratedCaller) GetPreviewBatchSwap(opts *bind.CallOpts, kind uint8, swaps []IERC20DexModuleBatchSwapStep) (struct {
	Asset  common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getPreviewBatchSwap", kind, swaps)

	outstruct := new(struct {
		Asset  common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Asset = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPreviewBatchSwap is a free data retrieval call binding the contract method 0xe6ee1e0b.
//
// Solidity: function getPreviewBatchSwap(uint8 kind, (address,address,uint256,address,uint256,bytes)[] swaps) view returns(address asset, uint256 amount)
func (_BexGenerated *BexGeneratedSession) GetPreviewBatchSwap(kind uint8, swaps []IERC20DexModuleBatchSwapStep) (struct {
	Asset  common.Address
	Amount *big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewBatchSwap(&_BexGenerated.CallOpts, kind, swaps)
}

// GetPreviewBatchSwap is a free data retrieval call binding the contract method 0xe6ee1e0b.
//
// Solidity: function getPreviewBatchSwap(uint8 kind, (address,address,uint256,address,uint256,bytes)[] swaps) view returns(address asset, uint256 amount)
func (_BexGenerated *BexGeneratedCallerSession) GetPreviewBatchSwap(kind uint8, swaps []IERC20DexModuleBatchSwapStep) (struct {
	Asset  common.Address
	Amount *big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewBatchSwap(&_BexGenerated.CallOpts, kind, swaps)
}

// GetPreviewBurnShares is a free data retrieval call binding the contract method 0x8dff6254.
//
// Solidity: function getPreviewBurnShares(address pool, address asset, uint256 amount) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCaller) GetPreviewBurnShares(opts *bind.CallOpts, pool common.Address, asset common.Address, amount *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getPreviewBurnShares", pool, asset, amount)

	outstruct := new(struct {
		Assets  []common.Address
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPreviewBurnShares is a free data retrieval call binding the contract method 0x8dff6254.
//
// Solidity: function getPreviewBurnShares(address pool, address asset, uint256 amount) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedSession) GetPreviewBurnShares(pool common.Address, asset common.Address, amount *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewBurnShares(&_BexGenerated.CallOpts, pool, asset, amount)
}

// GetPreviewBurnShares is a free data retrieval call binding the contract method 0x8dff6254.
//
// Solidity: function getPreviewBurnShares(address pool, address asset, uint256 amount) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCallerSession) GetPreviewBurnShares(pool common.Address, asset common.Address, amount *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewBurnShares(&_BexGenerated.CallOpts, pool, asset, amount)
}

// GetPreviewSharesForLiquidity is a free data retrieval call binding the contract method 0x6eb736db.
//
// Solidity: function getPreviewSharesForLiquidity(address pool, address[] assets, uint256[] amounts) view returns(address[] shares, uint256[] shareAmounts, address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedCaller) GetPreviewSharesForLiquidity(opts *bind.CallOpts, pool common.Address, assets []common.Address, amounts []*big.Int) (struct {
	Shares           []common.Address
	ShareAmounts     []*big.Int
	Liquidity        []common.Address
	LiquidityAmounts []*big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getPreviewSharesForLiquidity", pool, assets, amounts)

	outstruct := new(struct {
		Shares           []common.Address
		ShareAmounts     []*big.Int
		Liquidity        []common.Address
		LiquidityAmounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.ShareAmounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	outstruct.LiquidityAmounts = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPreviewSharesForLiquidity is a free data retrieval call binding the contract method 0x6eb736db.
//
// Solidity: function getPreviewSharesForLiquidity(address pool, address[] assets, uint256[] amounts) view returns(address[] shares, uint256[] shareAmounts, address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedSession) GetPreviewSharesForLiquidity(pool common.Address, assets []common.Address, amounts []*big.Int) (struct {
	Shares           []common.Address
	ShareAmounts     []*big.Int
	Liquidity        []common.Address
	LiquidityAmounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewSharesForLiquidity(&_BexGenerated.CallOpts, pool, assets, amounts)
}

// GetPreviewSharesForLiquidity is a free data retrieval call binding the contract method 0x6eb736db.
//
// Solidity: function getPreviewSharesForLiquidity(address pool, address[] assets, uint256[] amounts) view returns(address[] shares, uint256[] shareAmounts, address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedCallerSession) GetPreviewSharesForLiquidity(pool common.Address, assets []common.Address, amounts []*big.Int) (struct {
	Shares           []common.Address
	ShareAmounts     []*big.Int
	Liquidity        []common.Address
	LiquidityAmounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewSharesForLiquidity(&_BexGenerated.CallOpts, pool, assets, amounts)
}

// GetPreviewSharesForSingleSidedLiquidityRequest is a free data retrieval call binding the contract method 0x49ce7364.
//
// Solidity: function getPreviewSharesForSingleSidedLiquidityRequest(address pool, address asset, uint256 amount) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCaller) GetPreviewSharesForSingleSidedLiquidityRequest(opts *bind.CallOpts, pool common.Address, asset common.Address, amount *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getPreviewSharesForSingleSidedLiquidityRequest", pool, asset, amount)

	outstruct := new(struct {
		Assets  []common.Address
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPreviewSharesForSingleSidedLiquidityRequest is a free data retrieval call binding the contract method 0x49ce7364.
//
// Solidity: function getPreviewSharesForSingleSidedLiquidityRequest(address pool, address asset, uint256 amount) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedSession) GetPreviewSharesForSingleSidedLiquidityRequest(pool common.Address, asset common.Address, amount *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewSharesForSingleSidedLiquidityRequest(&_BexGenerated.CallOpts, pool, asset, amount)
}

// GetPreviewSharesForSingleSidedLiquidityRequest is a free data retrieval call binding the contract method 0x49ce7364.
//
// Solidity: function getPreviewSharesForSingleSidedLiquidityRequest(address pool, address asset, uint256 amount) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCallerSession) GetPreviewSharesForSingleSidedLiquidityRequest(pool common.Address, asset common.Address, amount *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewSharesForSingleSidedLiquidityRequest(&_BexGenerated.CallOpts, pool, asset, amount)
}

// GetPreviewSwapExact is a free data retrieval call binding the contract method 0xe2c2d064.
//
// Solidity: function getPreviewSwapExact(uint8 kind, address pool, address baseAsset, uint256 baseAssetAmount, address quoteAsset) view returns(address asset, uint256 amount)
func (_BexGenerated *BexGeneratedCaller) GetPreviewSwapExact(opts *bind.CallOpts, kind uint8, pool common.Address, baseAsset common.Address, baseAssetAmount *big.Int, quoteAsset common.Address) (struct {
	Asset  common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getPreviewSwapExact", kind, pool, baseAsset, baseAssetAmount, quoteAsset)

	outstruct := new(struct {
		Asset  common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Asset = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPreviewSwapExact is a free data retrieval call binding the contract method 0xe2c2d064.
//
// Solidity: function getPreviewSwapExact(uint8 kind, address pool, address baseAsset, uint256 baseAssetAmount, address quoteAsset) view returns(address asset, uint256 amount)
func (_BexGenerated *BexGeneratedSession) GetPreviewSwapExact(kind uint8, pool common.Address, baseAsset common.Address, baseAssetAmount *big.Int, quoteAsset common.Address) (struct {
	Asset  common.Address
	Amount *big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewSwapExact(&_BexGenerated.CallOpts, kind, pool, baseAsset, baseAssetAmount, quoteAsset)
}

// GetPreviewSwapExact is a free data retrieval call binding the contract method 0xe2c2d064.
//
// Solidity: function getPreviewSwapExact(uint8 kind, address pool, address baseAsset, uint256 baseAssetAmount, address quoteAsset) view returns(address asset, uint256 amount)
func (_BexGenerated *BexGeneratedCallerSession) GetPreviewSwapExact(kind uint8, pool common.Address, baseAsset common.Address, baseAssetAmount *big.Int, quoteAsset common.Address) (struct {
	Asset  common.Address
	Amount *big.Int
}, error) {
	return _BexGenerated.Contract.GetPreviewSwapExact(&_BexGenerated.CallOpts, kind, pool, baseAsset, baseAssetAmount, quoteAsset)
}

// GetRemoveLiquidityExactAmountOut is a free data retrieval call binding the contract method 0x46be96bb.
//
// Solidity: function getRemoveLiquidityExactAmountOut(address pool, address assetIn, uint256 assetAmount) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCaller) GetRemoveLiquidityExactAmountOut(opts *bind.CallOpts, pool common.Address, assetIn common.Address, assetAmount *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getRemoveLiquidityExactAmountOut", pool, assetIn, assetAmount)

	outstruct := new(struct {
		Assets  []common.Address
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRemoveLiquidityExactAmountOut is a free data retrieval call binding the contract method 0x46be96bb.
//
// Solidity: function getRemoveLiquidityExactAmountOut(address pool, address assetIn, uint256 assetAmount) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedSession) GetRemoveLiquidityExactAmountOut(pool common.Address, assetIn common.Address, assetAmount *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetRemoveLiquidityExactAmountOut(&_BexGenerated.CallOpts, pool, assetIn, assetAmount)
}

// GetRemoveLiquidityExactAmountOut is a free data retrieval call binding the contract method 0x46be96bb.
//
// Solidity: function getRemoveLiquidityExactAmountOut(address pool, address assetIn, uint256 assetAmount) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCallerSession) GetRemoveLiquidityExactAmountOut(pool common.Address, assetIn common.Address, assetAmount *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetRemoveLiquidityExactAmountOut(&_BexGenerated.CallOpts, pool, assetIn, assetAmount)
}

// GetRemoveLiquidityOneSideOut is a free data retrieval call binding the contract method 0xb82fa8d9.
//
// Solidity: function getRemoveLiquidityOneSideOut(address pool, address assetOut, uint256 sharesIn) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCaller) GetRemoveLiquidityOneSideOut(opts *bind.CallOpts, pool common.Address, assetOut common.Address, sharesIn *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getRemoveLiquidityOneSideOut", pool, assetOut, sharesIn)

	outstruct := new(struct {
		Assets  []common.Address
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRemoveLiquidityOneSideOut is a free data retrieval call binding the contract method 0xb82fa8d9.
//
// Solidity: function getRemoveLiquidityOneSideOut(address pool, address assetOut, uint256 sharesIn) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedSession) GetRemoveLiquidityOneSideOut(pool common.Address, assetOut common.Address, sharesIn *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetRemoveLiquidityOneSideOut(&_BexGenerated.CallOpts, pool, assetOut, sharesIn)
}

// GetRemoveLiquidityOneSideOut is a free data retrieval call binding the contract method 0xb82fa8d9.
//
// Solidity: function getRemoveLiquidityOneSideOut(address pool, address assetOut, uint256 sharesIn) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCallerSession) GetRemoveLiquidityOneSideOut(pool common.Address, assetOut common.Address, sharesIn *big.Int) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetRemoveLiquidityOneSideOut(&_BexGenerated.CallOpts, pool, assetOut, sharesIn)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xf20d8634.
//
// Solidity: function getTotalShares(address pool) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCaller) GetTotalShares(opts *bind.CallOpts, pool common.Address) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _BexGenerated.contract.Call(opts, &out, "getTotalShares", pool)

	outstruct := new(struct {
		Assets  []common.Address
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xf20d8634.
//
// Solidity: function getTotalShares(address pool) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedSession) GetTotalShares(pool common.Address) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetTotalShares(&_BexGenerated.CallOpts, pool)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xf20d8634.
//
// Solidity: function getTotalShares(address pool) view returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedCallerSession) GetTotalShares(pool common.Address) (struct {
	Assets  []common.Address
	Amounts []*big.Int
}, error) {
	return _BexGenerated.Contract.GetTotalShares(&_BexGenerated.CallOpts, pool)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x6d517aab.
//
// Solidity: function addLiquidity(address pool, address receiver, address[] assetsIn, uint256[] amountsIn) payable returns(address[] shares, uint256[] shareAmounts, address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedTransactor) AddLiquidity(opts *bind.TransactOpts, pool common.Address, receiver common.Address, assetsIn []common.Address, amountsIn []*big.Int) (*types.Transaction, error) {
	return _BexGenerated.contract.Transact(opts, "addLiquidity", pool, receiver, assetsIn, amountsIn)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x6d517aab.
//
// Solidity: function addLiquidity(address pool, address receiver, address[] assetsIn, uint256[] amountsIn) payable returns(address[] shares, uint256[] shareAmounts, address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedSession) AddLiquidity(pool common.Address, receiver common.Address, assetsIn []common.Address, amountsIn []*big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.AddLiquidity(&_BexGenerated.TransactOpts, pool, receiver, assetsIn, amountsIn)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x6d517aab.
//
// Solidity: function addLiquidity(address pool, address receiver, address[] assetsIn, uint256[] amountsIn) payable returns(address[] shares, uint256[] shareAmounts, address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedTransactorSession) AddLiquidity(pool common.Address, receiver common.Address, assetsIn []common.Address, amountsIn []*big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.AddLiquidity(&_BexGenerated.TransactOpts, pool, receiver, assetsIn, amountsIn)
}

// BatchSwap is a paid mutator transaction binding the contract method 0xe3414c00.
//
// Solidity: function batchSwap(uint8 kind, (address,address,uint256,address,uint256,bytes)[] swaps, uint256 deadline) payable returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedTransactor) BatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IERC20DexModuleBatchSwapStep, deadline *big.Int) (*types.Transaction, error) {
	return _BexGenerated.contract.Transact(opts, "batchSwap", kind, swaps, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0xe3414c00.
//
// Solidity: function batchSwap(uint8 kind, (address,address,uint256,address,uint256,bytes)[] swaps, uint256 deadline) payable returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedSession) BatchSwap(kind uint8, swaps []IERC20DexModuleBatchSwapStep, deadline *big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.BatchSwap(&_BexGenerated.TransactOpts, kind, swaps, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0xe3414c00.
//
// Solidity: function batchSwap(uint8 kind, (address,address,uint256,address,uint256,bytes)[] swaps, uint256 deadline) payable returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedTransactorSession) BatchSwap(kind uint8, swaps []IERC20DexModuleBatchSwapStep, deadline *big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.BatchSwap(&_BexGenerated.TransactOpts, kind, swaps, deadline)
}

// CreatePool is a paid mutator transaction binding the contract method 0x112d3dae.
//
// Solidity: function createPool(string name, address[] assetsIn, uint256[] amountsIn, string poolType, ((address,uint256)[],uint256) options) payable returns(address)
func (_BexGenerated *BexGeneratedTransactor) CreatePool(opts *bind.TransactOpts, name string, assetsIn []common.Address, amountsIn []*big.Int, poolType string, options IERC20DexModulePoolOptions) (*types.Transaction, error) {
	return _BexGenerated.contract.Transact(opts, "createPool", name, assetsIn, amountsIn, poolType, options)
}

// CreatePool is a paid mutator transaction binding the contract method 0x112d3dae.
//
// Solidity: function createPool(string name, address[] assetsIn, uint256[] amountsIn, string poolType, ((address,uint256)[],uint256) options) payable returns(address)
func (_BexGenerated *BexGeneratedSession) CreatePool(name string, assetsIn []common.Address, amountsIn []*big.Int, poolType string, options IERC20DexModulePoolOptions) (*types.Transaction, error) {
	return _BexGenerated.Contract.CreatePool(&_BexGenerated.TransactOpts, name, assetsIn, amountsIn, poolType, options)
}

// CreatePool is a paid mutator transaction binding the contract method 0x112d3dae.
//
// Solidity: function createPool(string name, address[] assetsIn, uint256[] amountsIn, string poolType, ((address,uint256)[],uint256) options) payable returns(address)
func (_BexGenerated *BexGeneratedTransactorSession) CreatePool(name string, assetsIn []common.Address, amountsIn []*big.Int, poolType string, options IERC20DexModulePoolOptions) (*types.Transaction, error) {
	return _BexGenerated.Contract.CreatePool(&_BexGenerated.TransactOpts, name, assetsIn, amountsIn, poolType, options)
}

// RemoveLiquidityBurningShares is a paid mutator transaction binding the contract method 0xe9942531.
//
// Solidity: function removeLiquidityBurningShares(address pool, address withdrawAddress, address assetIn, uint256 amountIn) payable returns(address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedTransactor) RemoveLiquidityBurningShares(opts *bind.TransactOpts, pool common.Address, withdrawAddress common.Address, assetIn common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _BexGenerated.contract.Transact(opts, "removeLiquidityBurningShares", pool, withdrawAddress, assetIn, amountIn)
}

// RemoveLiquidityBurningShares is a paid mutator transaction binding the contract method 0xe9942531.
//
// Solidity: function removeLiquidityBurningShares(address pool, address withdrawAddress, address assetIn, uint256 amountIn) payable returns(address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedSession) RemoveLiquidityBurningShares(pool common.Address, withdrawAddress common.Address, assetIn common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.RemoveLiquidityBurningShares(&_BexGenerated.TransactOpts, pool, withdrawAddress, assetIn, amountIn)
}

// RemoveLiquidityBurningShares is a paid mutator transaction binding the contract method 0xe9942531.
//
// Solidity: function removeLiquidityBurningShares(address pool, address withdrawAddress, address assetIn, uint256 amountIn) payable returns(address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedTransactorSession) RemoveLiquidityBurningShares(pool common.Address, withdrawAddress common.Address, assetIn common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.RemoveLiquidityBurningShares(&_BexGenerated.TransactOpts, pool, withdrawAddress, assetIn, amountIn)
}

// RemoveLiquidityExactAmount is a paid mutator transaction binding the contract method 0x4f9a546f.
//
// Solidity: function removeLiquidityExactAmount(address pool, address withdrawAddress, address assetOut, uint256 amountOut, address sharesIn, uint256 maxSharesIn) payable returns(address[] shares, uint256[] shareAmounts, address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedTransactor) RemoveLiquidityExactAmount(opts *bind.TransactOpts, pool common.Address, withdrawAddress common.Address, assetOut common.Address, amountOut *big.Int, sharesIn common.Address, maxSharesIn *big.Int) (*types.Transaction, error) {
	return _BexGenerated.contract.Transact(opts, "removeLiquidityExactAmount", pool, withdrawAddress, assetOut, amountOut, sharesIn, maxSharesIn)
}

// RemoveLiquidityExactAmount is a paid mutator transaction binding the contract method 0x4f9a546f.
//
// Solidity: function removeLiquidityExactAmount(address pool, address withdrawAddress, address assetOut, uint256 amountOut, address sharesIn, uint256 maxSharesIn) payable returns(address[] shares, uint256[] shareAmounts, address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedSession) RemoveLiquidityExactAmount(pool common.Address, withdrawAddress common.Address, assetOut common.Address, amountOut *big.Int, sharesIn common.Address, maxSharesIn *big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.RemoveLiquidityExactAmount(&_BexGenerated.TransactOpts, pool, withdrawAddress, assetOut, amountOut, sharesIn, maxSharesIn)
}

// RemoveLiquidityExactAmount is a paid mutator transaction binding the contract method 0x4f9a546f.
//
// Solidity: function removeLiquidityExactAmount(address pool, address withdrawAddress, address assetOut, uint256 amountOut, address sharesIn, uint256 maxSharesIn) payable returns(address[] shares, uint256[] shareAmounts, address[] liquidity, uint256[] liquidityAmounts)
func (_BexGenerated *BexGeneratedTransactorSession) RemoveLiquidityExactAmount(pool common.Address, withdrawAddress common.Address, assetOut common.Address, amountOut *big.Int, sharesIn common.Address, maxSharesIn *big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.RemoveLiquidityExactAmount(&_BexGenerated.TransactOpts, pool, withdrawAddress, assetOut, amountOut, sharesIn, maxSharesIn)
}

// Swap is a paid mutator transaction binding the contract method 0xc879007f.
//
// Solidity: function swap(uint8 kind, address poolId, address assetIn, uint256 amountIn, address assetOut, uint256 amountOut, uint256 deadline) payable returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedTransactor) Swap(opts *bind.TransactOpts, kind uint8, poolId common.Address, assetIn common.Address, amountIn *big.Int, assetOut common.Address, amountOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BexGenerated.contract.Transact(opts, "swap", kind, poolId, assetIn, amountIn, assetOut, amountOut, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0xc879007f.
//
// Solidity: function swap(uint8 kind, address poolId, address assetIn, uint256 amountIn, address assetOut, uint256 amountOut, uint256 deadline) payable returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedSession) Swap(kind uint8, poolId common.Address, assetIn common.Address, amountIn *big.Int, assetOut common.Address, amountOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.Swap(&_BexGenerated.TransactOpts, kind, poolId, assetIn, amountIn, assetOut, amountOut, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0xc879007f.
//
// Solidity: function swap(uint8 kind, address poolId, address assetIn, uint256 amountIn, address assetOut, uint256 amountOut, uint256 deadline) payable returns(address[] assets, uint256[] amounts)
func (_BexGenerated *BexGeneratedTransactorSession) Swap(kind uint8, poolId common.Address, assetIn common.Address, amountIn *big.Int, assetOut common.Address, amountOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BexGenerated.Contract.Swap(&_BexGenerated.TransactOpts, kind, poolId, assetIn, amountIn, assetOut, amountOut, deadline)
}
