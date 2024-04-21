// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bgt_station

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

// CosmosPageRequest is an auto generated low-level Go binding around an user-defined struct.
type CosmosPageRequest struct {
	Key        string
	Offset     uint64
	Limit      uint64
	CountTotal bool
	Reverse    bool
}

// CosmosPageResponse is an auto generated low-level Go binding around an user-defined struct.
type CosmosPageResponse struct {
	NextKey string
	Total   uint64
}

// IStakingModuleCommission is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleCommission struct {
	CommissionRates IStakingModuleCommissionRates
	UpdateTime      string
}

// IStakingModuleCommissionRates is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleCommissionRates struct {
	Rate          *big.Int
	MaxRate       *big.Int
	MaxChangeRate *big.Int
}

// IStakingModuleDelegation is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleDelegation struct {
	Delegator common.Address
	Balance   *big.Int
	Shares    *big.Int
}

// IStakingModuleDescription is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleDescription struct {
	Moniker         string
	Identity        string
	Website         string
	SecurityContact string
	Details         string
}

// IStakingModuleRedelegationEntry is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleRedelegationEntry struct {
	CreationHeight int64
	CompletionTime string
	InitialBalance *big.Int
	SharesDst      *big.Int
	UnbondingId    uint64
}

// IStakingModuleUnbondingDelegation is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleUnbondingDelegation struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Entries          []IStakingModuleUnbondingDelegationEntry
}

// IStakingModuleUnbondingDelegationEntry is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleUnbondingDelegationEntry struct {
	CreationHeight int64
	CompletionTime string
	InitialBalance *big.Int
	Balance        *big.Int
	UnbondingId    uint64
}

// IStakingModuleValidator is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleValidator struct {
	OperatorAddr            common.Address
	ConsAddr                []byte
	Jailed                  bool
	Status                  string
	Tokens                  *big.Int
	DelegatorShares         *big.Int
	Description             IStakingModuleDescription
	UnbondingHeight         int64
	UnbondingTime           string
	Commission              IStakingModuleCommission
	MinSelfDelegation       *big.Int
	UnbondingOnHoldRefCount int64
	UnbondingIds            []uint64
}

// StakingGeneratedModuleMetaData contains all meta data concerning the StakingGeneratedModule contract.
var StakingGeneratedModuleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"beginRedelegate\",\"inputs\":[{\"name\":\"srcValidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dstValidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelUnbondingDelegation\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creationHeight\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getBondedValidators\",\"inputs\":[{\"name\":\"pagination\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageRequest\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"countTotal\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reverse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingModule.Validator[]\",\"components\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"jailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"status\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegatorShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"securityContact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"unbondingHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"unbondingTime\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.Commission\",\"components\":[{\"name\":\"commissionRates\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.CommissionRates\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"updateTime\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"minSelfDelegation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"unbondingIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageResponse\",\"components\":[{\"name\":\"nextKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"total\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondedValidatorsByPower\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegation\",\"inputs\":[{\"name\":\"delegatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatorUnbondingDelegations\",\"inputs\":[{\"name\":\"delegatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pagination\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageRequest\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"countTotal\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reverse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingModule.UnbondingDelegation[]\",\"components\":[{\"name\":\"delegatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entries\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingModule.UnbondingDelegationEntry[]\",\"components\":[{\"name\":\"creationHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"completionTime\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondingId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageResponse\",\"components\":[{\"name\":\"nextKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"total\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatorValidators\",\"inputs\":[{\"name\":\"delegatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pagination\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageRequest\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"countTotal\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reverse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingModule.Validator[]\",\"components\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"jailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"status\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegatorShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"securityContact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"unbondingHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"unbondingTime\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.Commission\",\"components\":[{\"name\":\"commissionRates\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.CommissionRates\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"updateTime\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"minSelfDelegation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"unbondingIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageResponse\",\"components\":[{\"name\":\"nextKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"total\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedelegations\",\"inputs\":[{\"name\":\"delegatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"srcValidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dstValidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pagination\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageRequest\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"countTotal\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reverse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingModule.RedelegationEntry[]\",\"components\":[{\"name\":\"creationHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"completionTime\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sharesDst\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondingId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageResponse\",\"components\":[{\"name\":\"nextKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"total\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnbondingDelegation\",\"inputs\":[{\"name\":\"delegatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingModule.UnbondingDelegationEntry[]\",\"components\":[{\"name\":\"creationHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"completionTime\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondingId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValAddressFromConsAddress\",\"inputs\":[{\"name\":\"consAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.Validator\",\"components\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"jailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"status\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegatorShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"securityContact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"unbondingHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"unbondingTime\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.Commission\",\"components\":[{\"name\":\"commissionRates\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.CommissionRates\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"updateTime\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"minSelfDelegation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"unbondingIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDelegations\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pagination\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageRequest\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"countTotal\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reverse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingModule.Delegation[]\",\"components\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageResponse\",\"components\":[{\"name\":\"nextKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"total\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[{\"name\":\"pagination\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageRequest\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"countTotal\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reverse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingModule.Validator[]\",\"components\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"jailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"status\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegatorShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"securityContact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"unbondingHeight\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"unbondingTime\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.Commission\",\"components\":[{\"name\":\"commissionRates\",\"type\":\"tuple\",\"internalType\":\"structIStakingModule.CommissionRates\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"updateTime\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"minSelfDelegation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"unbondingIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageResponse\",\"components\":[{\"name\":\"nextKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"total\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"CancelUnbondingDelegation\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"creationHeight\",\"type\":\"int64\",\"indexed\":false,\"internalType\":\"int64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delegate\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Redelegate\",\"inputs\":[{\"name\":\"sourceValidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destinationValidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unbond\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false}]",
}

// StakingGeneratedModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingGeneratedModuleMetaData.ABI instead.
var StakingGeneratedModuleABI = StakingGeneratedModuleMetaData.ABI

// StakingGeneratedModule is an auto generated Go binding around an Ethereum contract.
type StakingGeneratedModule struct {
	StakingGeneratedModuleCaller     // Read-only binding to the contract
	StakingGeneratedModuleTransactor // Write-only binding to the contract
	StakingGeneratedModuleFilterer   // Log filterer for contract events
}

// StakingGeneratedModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingGeneratedModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingGeneratedModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingGeneratedModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingGeneratedModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingGeneratedModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingGeneratedModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingGeneratedModuleSession struct {
	Contract     *StakingGeneratedModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakingGeneratedModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingGeneratedModuleCallerSession struct {
	Contract *StakingGeneratedModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StakingGeneratedModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingGeneratedModuleTransactorSession struct {
	Contract     *StakingGeneratedModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StakingGeneratedModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingGeneratedModuleRaw struct {
	Contract *StakingGeneratedModule // Generic contract binding to access the raw methods on
}

// StakingGeneratedModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingGeneratedModuleCallerRaw struct {
	Contract *StakingGeneratedModuleCaller // Generic read-only contract binding to access the raw methods on
}

// StakingGeneratedModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingGeneratedModuleTransactorRaw struct {
	Contract *StakingGeneratedModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingGeneratedModule creates a new instance of StakingGeneratedModule, bound to a specific deployed contract.
func NewStakingGeneratedModule(address common.Address, backend bind.ContractBackend) (*StakingGeneratedModule, error) {
	contract, err := bindStakingGeneratedModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingGeneratedModule{StakingGeneratedModuleCaller: StakingGeneratedModuleCaller{contract: contract}, StakingGeneratedModuleTransactor: StakingGeneratedModuleTransactor{contract: contract}, StakingGeneratedModuleFilterer: StakingGeneratedModuleFilterer{contract: contract}}, nil
}

// NewStakingGeneratedModuleCaller creates a new read-only instance of StakingGeneratedModule, bound to a specific deployed contract.
func NewStakingGeneratedModuleCaller(address common.Address, caller bind.ContractCaller) (*StakingGeneratedModuleCaller, error) {
	contract, err := bindStakingGeneratedModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingGeneratedModuleCaller{contract: contract}, nil
}

// NewStakingGeneratedModuleTransactor creates a new write-only instance of StakingGeneratedModule, bound to a specific deployed contract.
func NewStakingGeneratedModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingGeneratedModuleTransactor, error) {
	contract, err := bindStakingGeneratedModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingGeneratedModuleTransactor{contract: contract}, nil
}

// NewStakingGeneratedModuleFilterer creates a new log filterer instance of StakingGeneratedModule, bound to a specific deployed contract.
func NewStakingGeneratedModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingGeneratedModuleFilterer, error) {
	contract, err := bindStakingGeneratedModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingGeneratedModuleFilterer{contract: contract}, nil
}

// bindStakingGeneratedModule binds a generic wrapper to an already deployed contract.
func bindStakingGeneratedModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingGeneratedModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingGeneratedModule *StakingGeneratedModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingGeneratedModule.Contract.StakingGeneratedModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingGeneratedModule *StakingGeneratedModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.StakingGeneratedModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingGeneratedModule *StakingGeneratedModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.StakingGeneratedModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingGeneratedModule *StakingGeneratedModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingGeneratedModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingGeneratedModule *StakingGeneratedModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingGeneratedModule *StakingGeneratedModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.contract.Transact(opts, method, params...)
}

// GetBondedValidators is a free data retrieval call binding the contract method 0xcf3f2340.
//
// Solidity: function getBondedValidators((string,uint64,uint64,bool,bool) pagination) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetBondedValidators(opts *bind.CallOpts, pagination CosmosPageRequest) ([]IStakingModuleValidator, CosmosPageResponse, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getBondedValidators", pagination)

	if err != nil {
		return *new([]IStakingModuleValidator), *new(CosmosPageResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleValidator)).(*[]IStakingModuleValidator)
	out1 := *abi.ConvertType(out[1], new(CosmosPageResponse)).(*CosmosPageResponse)

	return out0, out1, err

}

// GetBondedValidators is a free data retrieval call binding the contract method 0xcf3f2340.
//
// Solidity: function getBondedValidators((string,uint64,uint64,bool,bool) pagination) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetBondedValidators(pagination CosmosPageRequest) ([]IStakingModuleValidator, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetBondedValidators(&_StakingGeneratedModule.CallOpts, pagination)
}

// GetBondedValidators is a free data retrieval call binding the contract method 0xcf3f2340.
//
// Solidity: function getBondedValidators((string,uint64,uint64,bool,bool) pagination) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetBondedValidators(pagination CosmosPageRequest) ([]IStakingModuleValidator, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetBondedValidators(&_StakingGeneratedModule.CallOpts, pagination)
}

// GetBondedValidatorsByPower is a free data retrieval call binding the contract method 0xdcaf464a.
//
// Solidity: function getBondedValidatorsByPower() view returns(address[])
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetBondedValidatorsByPower(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getBondedValidatorsByPower")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBondedValidatorsByPower is a free data retrieval call binding the contract method 0xdcaf464a.
//
// Solidity: function getBondedValidatorsByPower() view returns(address[])
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetBondedValidatorsByPower() ([]common.Address, error) {
	return _StakingGeneratedModule.Contract.GetBondedValidatorsByPower(&_StakingGeneratedModule.CallOpts)
}

// GetBondedValidatorsByPower is a free data retrieval call binding the contract method 0xdcaf464a.
//
// Solidity: function getBondedValidatorsByPower() view returns(address[])
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetBondedValidatorsByPower() ([]common.Address, error) {
	return _StakingGeneratedModule.Contract.GetBondedValidatorsByPower(&_StakingGeneratedModule.CallOpts)
}

// GetDelegation is a free data retrieval call binding the contract method 0x15049a5a.
//
// Solidity: function getDelegation(address delegatorAddress, address validatorAddress) view returns(uint256)
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetDelegation(opts *bind.CallOpts, delegatorAddress common.Address, validatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getDelegation", delegatorAddress, validatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegation is a free data retrieval call binding the contract method 0x15049a5a.
//
// Solidity: function getDelegation(address delegatorAddress, address validatorAddress) view returns(uint256)
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetDelegation(delegatorAddress common.Address, validatorAddress common.Address) (*big.Int, error) {
	return _StakingGeneratedModule.Contract.GetDelegation(&_StakingGeneratedModule.CallOpts, delegatorAddress, validatorAddress)
}

// GetDelegation is a free data retrieval call binding the contract method 0x15049a5a.
//
// Solidity: function getDelegation(address delegatorAddress, address validatorAddress) view returns(uint256)
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetDelegation(delegatorAddress common.Address, validatorAddress common.Address) (*big.Int, error) {
	return _StakingGeneratedModule.Contract.GetDelegation(&_StakingGeneratedModule.CallOpts, delegatorAddress, validatorAddress)
}

// GetDelegatorUnbondingDelegations is a free data retrieval call binding the contract method 0xd2b3c8fe.
//
// Solidity: function getDelegatorUnbondingDelegations(address delegatorAddress, (string,uint64,uint64,bool,bool) pagination) view returns((address,address,(int64,string,uint256,uint256,uint64)[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetDelegatorUnbondingDelegations(opts *bind.CallOpts, delegatorAddress common.Address, pagination CosmosPageRequest) ([]IStakingModuleUnbondingDelegation, CosmosPageResponse, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getDelegatorUnbondingDelegations", delegatorAddress, pagination)

	if err != nil {
		return *new([]IStakingModuleUnbondingDelegation), *new(CosmosPageResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleUnbondingDelegation)).(*[]IStakingModuleUnbondingDelegation)
	out1 := *abi.ConvertType(out[1], new(CosmosPageResponse)).(*CosmosPageResponse)

	return out0, out1, err

}

// GetDelegatorUnbondingDelegations is a free data retrieval call binding the contract method 0xd2b3c8fe.
//
// Solidity: function getDelegatorUnbondingDelegations(address delegatorAddress, (string,uint64,uint64,bool,bool) pagination) view returns((address,address,(int64,string,uint256,uint256,uint64)[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetDelegatorUnbondingDelegations(delegatorAddress common.Address, pagination CosmosPageRequest) ([]IStakingModuleUnbondingDelegation, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetDelegatorUnbondingDelegations(&_StakingGeneratedModule.CallOpts, delegatorAddress, pagination)
}

// GetDelegatorUnbondingDelegations is a free data retrieval call binding the contract method 0xd2b3c8fe.
//
// Solidity: function getDelegatorUnbondingDelegations(address delegatorAddress, (string,uint64,uint64,bool,bool) pagination) view returns((address,address,(int64,string,uint256,uint256,uint64)[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetDelegatorUnbondingDelegations(delegatorAddress common.Address, pagination CosmosPageRequest) ([]IStakingModuleUnbondingDelegation, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetDelegatorUnbondingDelegations(&_StakingGeneratedModule.CallOpts, delegatorAddress, pagination)
}

// GetDelegatorValidators is a free data retrieval call binding the contract method 0xb067c623.
//
// Solidity: function getDelegatorValidators(address delegatorAddress, (string,uint64,uint64,bool,bool) pagination) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetDelegatorValidators(opts *bind.CallOpts, delegatorAddress common.Address, pagination CosmosPageRequest) ([]IStakingModuleValidator, CosmosPageResponse, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getDelegatorValidators", delegatorAddress, pagination)

	if err != nil {
		return *new([]IStakingModuleValidator), *new(CosmosPageResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleValidator)).(*[]IStakingModuleValidator)
	out1 := *abi.ConvertType(out[1], new(CosmosPageResponse)).(*CosmosPageResponse)

	return out0, out1, err

}

// GetDelegatorValidators is a free data retrieval call binding the contract method 0xb067c623.
//
// Solidity: function getDelegatorValidators(address delegatorAddress, (string,uint64,uint64,bool,bool) pagination) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetDelegatorValidators(delegatorAddress common.Address, pagination CosmosPageRequest) ([]IStakingModuleValidator, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetDelegatorValidators(&_StakingGeneratedModule.CallOpts, delegatorAddress, pagination)
}

// GetDelegatorValidators is a free data retrieval call binding the contract method 0xb067c623.
//
// Solidity: function getDelegatorValidators(address delegatorAddress, (string,uint64,uint64,bool,bool) pagination) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetDelegatorValidators(delegatorAddress common.Address, pagination CosmosPageRequest) ([]IStakingModuleValidator, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetDelegatorValidators(&_StakingGeneratedModule.CallOpts, delegatorAddress, pagination)
}

// GetRedelegations is a free data retrieval call binding the contract method 0x1c441040.
//
// Solidity: function getRedelegations(address delegatorAddress, address srcValidator, address dstValidator, (string,uint64,uint64,bool,bool) pagination) view returns((int64,string,uint256,uint256,uint64)[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetRedelegations(opts *bind.CallOpts, delegatorAddress common.Address, srcValidator common.Address, dstValidator common.Address, pagination CosmosPageRequest) ([]IStakingModuleRedelegationEntry, CosmosPageResponse, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getRedelegations", delegatorAddress, srcValidator, dstValidator, pagination)

	if err != nil {
		return *new([]IStakingModuleRedelegationEntry), *new(CosmosPageResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleRedelegationEntry)).(*[]IStakingModuleRedelegationEntry)
	out1 := *abi.ConvertType(out[1], new(CosmosPageResponse)).(*CosmosPageResponse)

	return out0, out1, err

}

// GetRedelegations is a free data retrieval call binding the contract method 0x1c441040.
//
// Solidity: function getRedelegations(address delegatorAddress, address srcValidator, address dstValidator, (string,uint64,uint64,bool,bool) pagination) view returns((int64,string,uint256,uint256,uint64)[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetRedelegations(delegatorAddress common.Address, srcValidator common.Address, dstValidator common.Address, pagination CosmosPageRequest) ([]IStakingModuleRedelegationEntry, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetRedelegations(&_StakingGeneratedModule.CallOpts, delegatorAddress, srcValidator, dstValidator, pagination)
}

// GetRedelegations is a free data retrieval call binding the contract method 0x1c441040.
//
// Solidity: function getRedelegations(address delegatorAddress, address srcValidator, address dstValidator, (string,uint64,uint64,bool,bool) pagination) view returns((int64,string,uint256,uint256,uint64)[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetRedelegations(delegatorAddress common.Address, srcValidator common.Address, dstValidator common.Address, pagination CosmosPageRequest) ([]IStakingModuleRedelegationEntry, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetRedelegations(&_StakingGeneratedModule.CallOpts, delegatorAddress, srcValidator, dstValidator, pagination)
}

// GetUnbondingDelegation is a free data retrieval call binding the contract method 0xc60b8213.
//
// Solidity: function getUnbondingDelegation(address delegatorAddress, address validatorAddress) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetUnbondingDelegation(opts *bind.CallOpts, delegatorAddress common.Address, validatorAddress common.Address) ([]IStakingModuleUnbondingDelegationEntry, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getUnbondingDelegation", delegatorAddress, validatorAddress)

	if err != nil {
		return *new([]IStakingModuleUnbondingDelegationEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleUnbondingDelegationEntry)).(*[]IStakingModuleUnbondingDelegationEntry)

	return out0, err

}

// GetUnbondingDelegation is a free data retrieval call binding the contract method 0xc60b8213.
//
// Solidity: function getUnbondingDelegation(address delegatorAddress, address validatorAddress) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetUnbondingDelegation(delegatorAddress common.Address, validatorAddress common.Address) ([]IStakingModuleUnbondingDelegationEntry, error) {
	return _StakingGeneratedModule.Contract.GetUnbondingDelegation(&_StakingGeneratedModule.CallOpts, delegatorAddress, validatorAddress)
}

// GetUnbondingDelegation is a free data retrieval call binding the contract method 0xc60b8213.
//
// Solidity: function getUnbondingDelegation(address delegatorAddress, address validatorAddress) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetUnbondingDelegation(delegatorAddress common.Address, validatorAddress common.Address) ([]IStakingModuleUnbondingDelegationEntry, error) {
	return _StakingGeneratedModule.Contract.GetUnbondingDelegation(&_StakingGeneratedModule.CallOpts, delegatorAddress, validatorAddress)
}

// GetValAddressFromConsAddress is a free data retrieval call binding the contract method 0x411dee2d.
//
// Solidity: function getValAddressFromConsAddress(bytes consAddr) pure returns(address)
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetValAddressFromConsAddress(opts *bind.CallOpts, consAddr []byte) (common.Address, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getValAddressFromConsAddress", consAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValAddressFromConsAddress is a free data retrieval call binding the contract method 0x411dee2d.
//
// Solidity: function getValAddressFromConsAddress(bytes consAddr) pure returns(address)
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetValAddressFromConsAddress(consAddr []byte) (common.Address, error) {
	return _StakingGeneratedModule.Contract.GetValAddressFromConsAddress(&_StakingGeneratedModule.CallOpts, consAddr)
}

// GetValAddressFromConsAddress is a free data retrieval call binding the contract method 0x411dee2d.
//
// Solidity: function getValAddressFromConsAddress(bytes consAddr) pure returns(address)
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetValAddressFromConsAddress(consAddr []byte) (common.Address, error) {
	return _StakingGeneratedModule.Contract.GetValAddressFromConsAddress(&_StakingGeneratedModule.CallOpts, consAddr)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validatorAddress) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[]))
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetValidator(opts *bind.CallOpts, validatorAddress common.Address) (IStakingModuleValidator, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getValidator", validatorAddress)

	if err != nil {
		return *new(IStakingModuleValidator), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakingModuleValidator)).(*IStakingModuleValidator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validatorAddress) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[]))
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetValidator(validatorAddress common.Address) (IStakingModuleValidator, error) {
	return _StakingGeneratedModule.Contract.GetValidator(&_StakingGeneratedModule.CallOpts, validatorAddress)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validatorAddress) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[]))
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetValidator(validatorAddress common.Address) (IStakingModuleValidator, error) {
	return _StakingGeneratedModule.Contract.GetValidator(&_StakingGeneratedModule.CallOpts, validatorAddress)
}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x1f360742.
//
// Solidity: function getValidatorDelegations(address validatorAddress, (string,uint64,uint64,bool,bool) pagination) view returns((address,uint256,uint256)[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetValidatorDelegations(opts *bind.CallOpts, validatorAddress common.Address, pagination CosmosPageRequest) ([]IStakingModuleDelegation, CosmosPageResponse, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getValidatorDelegations", validatorAddress, pagination)

	if err != nil {
		return *new([]IStakingModuleDelegation), *new(CosmosPageResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleDelegation)).(*[]IStakingModuleDelegation)
	out1 := *abi.ConvertType(out[1], new(CosmosPageResponse)).(*CosmosPageResponse)

	return out0, out1, err

}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x1f360742.
//
// Solidity: function getValidatorDelegations(address validatorAddress, (string,uint64,uint64,bool,bool) pagination) view returns((address,uint256,uint256)[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetValidatorDelegations(validatorAddress common.Address, pagination CosmosPageRequest) ([]IStakingModuleDelegation, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetValidatorDelegations(&_StakingGeneratedModule.CallOpts, validatorAddress, pagination)
}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x1f360742.
//
// Solidity: function getValidatorDelegations(address validatorAddress, (string,uint64,uint64,bool,bool) pagination) view returns((address,uint256,uint256)[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetValidatorDelegations(validatorAddress common.Address, pagination CosmosPageRequest) ([]IStakingModuleDelegation, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetValidatorDelegations(&_StakingGeneratedModule.CallOpts, validatorAddress, pagination)
}

// GetValidators is a free data retrieval call binding the contract method 0xbfc4dcd5.
//
// Solidity: function getValidators((string,uint64,uint64,bool,bool) pagination) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCaller) GetValidators(opts *bind.CallOpts, pagination CosmosPageRequest) ([]IStakingModuleValidator, CosmosPageResponse, error) {
	var out []interface{}
	err := _StakingGeneratedModule.contract.Call(opts, &out, "getValidators", pagination)

	if err != nil {
		return *new([]IStakingModuleValidator), *new(CosmosPageResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleValidator)).(*[]IStakingModuleValidator)
	out1 := *abi.ConvertType(out[1], new(CosmosPageResponse)).(*CosmosPageResponse)

	return out0, out1, err

}

// GetValidators is a free data retrieval call binding the contract method 0xbfc4dcd5.
//
// Solidity: function getValidators((string,uint64,uint64,bool,bool) pagination) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleSession) GetValidators(pagination CosmosPageRequest) ([]IStakingModuleValidator, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetValidators(&_StakingGeneratedModule.CallOpts, pagination)
}

// GetValidators is a free data retrieval call binding the contract method 0xbfc4dcd5.
//
// Solidity: function getValidators((string,uint64,uint64,bool,bool) pagination) view returns((address,bytes,bool,string,uint256,uint256,(string,string,string,string,string),int64,string,((uint256,uint256,uint256),string),uint256,int64,uint64[])[], (string,uint64))
func (_StakingGeneratedModule *StakingGeneratedModuleCallerSession) GetValidators(pagination CosmosPageRequest) ([]IStakingModuleValidator, CosmosPageResponse, error) {
	return _StakingGeneratedModule.Contract.GetValidators(&_StakingGeneratedModule.CallOpts, pagination)
}

// BeginRedelegate is a paid mutator transaction binding the contract method 0xb3a8ae3b.
//
// Solidity: function beginRedelegate(address srcValidator, address dstValidator, uint256 amount) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleTransactor) BeginRedelegate(opts *bind.TransactOpts, srcValidator common.Address, dstValidator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingGeneratedModule.contract.Transact(opts, "beginRedelegate", srcValidator, dstValidator, amount)
}

// BeginRedelegate is a paid mutator transaction binding the contract method 0xb3a8ae3b.
//
// Solidity: function beginRedelegate(address srcValidator, address dstValidator, uint256 amount) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleSession) BeginRedelegate(srcValidator common.Address, dstValidator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.BeginRedelegate(&_StakingGeneratedModule.TransactOpts, srcValidator, dstValidator, amount)
}

// BeginRedelegate is a paid mutator transaction binding the contract method 0xb3a8ae3b.
//
// Solidity: function beginRedelegate(address srcValidator, address dstValidator, uint256 amount) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleTransactorSession) BeginRedelegate(srcValidator common.Address, dstValidator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.BeginRedelegate(&_StakingGeneratedModule.TransactOpts, srcValidator, dstValidator, amount)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x69a2f536.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, int64 creationHeight) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleTransactor) CancelUnbondingDelegation(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int, creationHeight int64) (*types.Transaction, error) {
	return _StakingGeneratedModule.contract.Transact(opts, "cancelUnbondingDelegation", validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x69a2f536.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, int64 creationHeight) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleSession) CancelUnbondingDelegation(validatorAddress common.Address, amount *big.Int, creationHeight int64) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.CancelUnbondingDelegation(&_StakingGeneratedModule.TransactOpts, validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x69a2f536.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, int64 creationHeight) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleTransactorSession) CancelUnbondingDelegation(validatorAddress common.Address, amount *big.Int, creationHeight int64) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.CancelUnbondingDelegation(&_StakingGeneratedModule.TransactOpts, validatorAddress, amount, creationHeight)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleTransactor) Delegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingGeneratedModule.contract.Transact(opts, "delegate", validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleSession) Delegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.Delegate(&_StakingGeneratedModule.TransactOpts, validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleTransactorSession) Delegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.Delegate(&_StakingGeneratedModule.TransactOpts, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleTransactor) Undelegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingGeneratedModule.contract.Transact(opts, "undelegate", validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.Undelegate(&_StakingGeneratedModule.TransactOpts, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) payable returns(bool)
func (_StakingGeneratedModule *StakingGeneratedModuleTransactorSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingGeneratedModule.Contract.Undelegate(&_StakingGeneratedModule.TransactOpts, validatorAddress, amount)
}

// StakingGeneratedModuleCancelUnbondingDelegationIterator is returned from FilterCancelUnbondingDelegation and is used to iterate over the raw logs and unpacked data for CancelUnbondingDelegation events raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleCancelUnbondingDelegationIterator struct {
	Event *StakingGeneratedModuleCancelUnbondingDelegation // Event containing the contract specifics and raw log

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
func (it *StakingGeneratedModuleCancelUnbondingDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingGeneratedModuleCancelUnbondingDelegation)
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
		it.Event = new(StakingGeneratedModuleCancelUnbondingDelegation)
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
func (it *StakingGeneratedModuleCancelUnbondingDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingGeneratedModuleCancelUnbondingDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingGeneratedModuleCancelUnbondingDelegation represents a CancelUnbondingDelegation event raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleCancelUnbondingDelegation struct {
	Validator      common.Address
	Delegator      common.Address
	Amount         []CosmosCoin
	CreationHeight int64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCancelUnbondingDelegation is a free log retrieval operation binding the contract event 0x30c2800a487f4891694e92c2748f62229fc352c93ae350a7ff63e3ab605a4aa5.
//
// Solidity: event CancelUnbondingDelegation(address indexed validator, address indexed delegator, (uint256,string)[] amount, int64 creationHeight)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) FilterCancelUnbondingDelegation(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*StakingGeneratedModuleCancelUnbondingDelegationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.FilterLogs(opts, "CancelUnbondingDelegation", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingGeneratedModuleCancelUnbondingDelegationIterator{contract: _StakingGeneratedModule.contract, event: "CancelUnbondingDelegation", logs: logs, sub: sub}, nil
}

// WatchCancelUnbondingDelegation is a free log subscription operation binding the contract event 0x30c2800a487f4891694e92c2748f62229fc352c93ae350a7ff63e3ab605a4aa5.
//
// Solidity: event CancelUnbondingDelegation(address indexed validator, address indexed delegator, (uint256,string)[] amount, int64 creationHeight)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) WatchCancelUnbondingDelegation(opts *bind.WatchOpts, sink chan<- *StakingGeneratedModuleCancelUnbondingDelegation, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.WatchLogs(opts, "CancelUnbondingDelegation", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingGeneratedModuleCancelUnbondingDelegation)
				if err := _StakingGeneratedModule.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
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

// ParseCancelUnbondingDelegation is a log parse operation binding the contract event 0x30c2800a487f4891694e92c2748f62229fc352c93ae350a7ff63e3ab605a4aa5.
//
// Solidity: event CancelUnbondingDelegation(address indexed validator, address indexed delegator, (uint256,string)[] amount, int64 creationHeight)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) ParseCancelUnbondingDelegation(log types.Log) (*StakingGeneratedModuleCancelUnbondingDelegation, error) {
	event := new(StakingGeneratedModuleCancelUnbondingDelegation)
	if err := _StakingGeneratedModule.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingGeneratedModuleCreateValidatorIterator is returned from FilterCreateValidator and is used to iterate over the raw logs and unpacked data for CreateValidator events raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleCreateValidatorIterator struct {
	Event *StakingGeneratedModuleCreateValidator // Event containing the contract specifics and raw log

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
func (it *StakingGeneratedModuleCreateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingGeneratedModuleCreateValidator)
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
		it.Event = new(StakingGeneratedModuleCreateValidator)
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
func (it *StakingGeneratedModuleCreateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingGeneratedModuleCreateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingGeneratedModuleCreateValidator represents a CreateValidator event raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleCreateValidator struct {
	Validator common.Address
	Amount    []CosmosCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateValidator is a free log retrieval operation binding the contract event 0x2bc39078c6a3394560520acda6eedb30ee0e6a2cf247ebf0857d3f3e11bd69e8.
//
// Solidity: event CreateValidator(address indexed validator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) FilterCreateValidator(opts *bind.FilterOpts, validator []common.Address) (*StakingGeneratedModuleCreateValidatorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.FilterLogs(opts, "CreateValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingGeneratedModuleCreateValidatorIterator{contract: _StakingGeneratedModule.contract, event: "CreateValidator", logs: logs, sub: sub}, nil
}

// WatchCreateValidator is a free log subscription operation binding the contract event 0x2bc39078c6a3394560520acda6eedb30ee0e6a2cf247ebf0857d3f3e11bd69e8.
//
// Solidity: event CreateValidator(address indexed validator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) WatchCreateValidator(opts *bind.WatchOpts, sink chan<- *StakingGeneratedModuleCreateValidator, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.WatchLogs(opts, "CreateValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingGeneratedModuleCreateValidator)
				if err := _StakingGeneratedModule.contract.UnpackLog(event, "CreateValidator", log); err != nil {
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

// ParseCreateValidator is a log parse operation binding the contract event 0x2bc39078c6a3394560520acda6eedb30ee0e6a2cf247ebf0857d3f3e11bd69e8.
//
// Solidity: event CreateValidator(address indexed validator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) ParseCreateValidator(log types.Log) (*StakingGeneratedModuleCreateValidator, error) {
	event := new(StakingGeneratedModuleCreateValidator)
	if err := _StakingGeneratedModule.contract.UnpackLog(event, "CreateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingGeneratedModuleDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleDelegateIterator struct {
	Event *StakingGeneratedModuleDelegate // Event containing the contract specifics and raw log

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
func (it *StakingGeneratedModuleDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingGeneratedModuleDelegate)
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
		it.Event = new(StakingGeneratedModuleDelegate)
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
func (it *StakingGeneratedModuleDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingGeneratedModuleDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingGeneratedModuleDelegate represents a Delegate event raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleDelegate struct {
	Validator common.Address
	Amount    []CosmosCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0x86d06596b16cc784c8990ddf4c3e195fd968c238f5999435057d48c77ba49f2f.
//
// Solidity: event Delegate(address indexed validator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) FilterDelegate(opts *bind.FilterOpts, validator []common.Address) (*StakingGeneratedModuleDelegateIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.FilterLogs(opts, "Delegate", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingGeneratedModuleDelegateIterator{contract: _StakingGeneratedModule.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0x86d06596b16cc784c8990ddf4c3e195fd968c238f5999435057d48c77ba49f2f.
//
// Solidity: event Delegate(address indexed validator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *StakingGeneratedModuleDelegate, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.WatchLogs(opts, "Delegate", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingGeneratedModuleDelegate)
				if err := _StakingGeneratedModule.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0x86d06596b16cc784c8990ddf4c3e195fd968c238f5999435057d48c77ba49f2f.
//
// Solidity: event Delegate(address indexed validator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) ParseDelegate(log types.Log) (*StakingGeneratedModuleDelegate, error) {
	event := new(StakingGeneratedModuleDelegate)
	if err := _StakingGeneratedModule.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingGeneratedModuleRedelegateIterator is returned from FilterRedelegate and is used to iterate over the raw logs and unpacked data for Redelegate events raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleRedelegateIterator struct {
	Event *StakingGeneratedModuleRedelegate // Event containing the contract specifics and raw log

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
func (it *StakingGeneratedModuleRedelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingGeneratedModuleRedelegate)
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
		it.Event = new(StakingGeneratedModuleRedelegate)
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
func (it *StakingGeneratedModuleRedelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingGeneratedModuleRedelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingGeneratedModuleRedelegate represents a Redelegate event raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleRedelegate struct {
	SourceValidator      common.Address
	DestinationValidator common.Address
	Amount               []CosmosCoin
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRedelegate is a free log retrieval operation binding the contract event 0xe723c90c78f428142cef6e47c9395b54bab8137eaaa44f34a1edbf930114c1eb.
//
// Solidity: event Redelegate(address indexed sourceValidator, address indexed destinationValidator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) FilterRedelegate(opts *bind.FilterOpts, sourceValidator []common.Address, destinationValidator []common.Address) (*StakingGeneratedModuleRedelegateIterator, error) {

	var sourceValidatorRule []interface{}
	for _, sourceValidatorItem := range sourceValidator {
		sourceValidatorRule = append(sourceValidatorRule, sourceValidatorItem)
	}
	var destinationValidatorRule []interface{}
	for _, destinationValidatorItem := range destinationValidator {
		destinationValidatorRule = append(destinationValidatorRule, destinationValidatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.FilterLogs(opts, "Redelegate", sourceValidatorRule, destinationValidatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingGeneratedModuleRedelegateIterator{contract: _StakingGeneratedModule.contract, event: "Redelegate", logs: logs, sub: sub}, nil
}

// WatchRedelegate is a free log subscription operation binding the contract event 0xe723c90c78f428142cef6e47c9395b54bab8137eaaa44f34a1edbf930114c1eb.
//
// Solidity: event Redelegate(address indexed sourceValidator, address indexed destinationValidator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) WatchRedelegate(opts *bind.WatchOpts, sink chan<- *StakingGeneratedModuleRedelegate, sourceValidator []common.Address, destinationValidator []common.Address) (event.Subscription, error) {

	var sourceValidatorRule []interface{}
	for _, sourceValidatorItem := range sourceValidator {
		sourceValidatorRule = append(sourceValidatorRule, sourceValidatorItem)
	}
	var destinationValidatorRule []interface{}
	for _, destinationValidatorItem := range destinationValidator {
		destinationValidatorRule = append(destinationValidatorRule, destinationValidatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.WatchLogs(opts, "Redelegate", sourceValidatorRule, destinationValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingGeneratedModuleRedelegate)
				if err := _StakingGeneratedModule.contract.UnpackLog(event, "Redelegate", log); err != nil {
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

// ParseRedelegate is a log parse operation binding the contract event 0xe723c90c78f428142cef6e47c9395b54bab8137eaaa44f34a1edbf930114c1eb.
//
// Solidity: event Redelegate(address indexed sourceValidator, address indexed destinationValidator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) ParseRedelegate(log types.Log) (*StakingGeneratedModuleRedelegate, error) {
	event := new(StakingGeneratedModuleRedelegate)
	if err := _StakingGeneratedModule.contract.UnpackLog(event, "Redelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingGeneratedModuleUnbondIterator is returned from FilterUnbond and is used to iterate over the raw logs and unpacked data for Unbond events raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleUnbondIterator struct {
	Event *StakingGeneratedModuleUnbond // Event containing the contract specifics and raw log

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
func (it *StakingGeneratedModuleUnbondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingGeneratedModuleUnbond)
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
		it.Event = new(StakingGeneratedModuleUnbond)
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
func (it *StakingGeneratedModuleUnbondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingGeneratedModuleUnbondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingGeneratedModuleUnbond represents a Unbond event raised by the StakingGeneratedModule contract.
type StakingGeneratedModuleUnbond struct {
	Validator common.Address
	Amount    []CosmosCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnbond is a free log retrieval operation binding the contract event 0x9b3dc86ff4188cb66e4fbacecb81f0fa1648e8fde176887bdfedafb075f5bb3e.
//
// Solidity: event Unbond(address indexed validator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) FilterUnbond(opts *bind.FilterOpts, validator []common.Address) (*StakingGeneratedModuleUnbondIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.FilterLogs(opts, "Unbond", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingGeneratedModuleUnbondIterator{contract: _StakingGeneratedModule.contract, event: "Unbond", logs: logs, sub: sub}, nil
}

// WatchUnbond is a free log subscription operation binding the contract event 0x9b3dc86ff4188cb66e4fbacecb81f0fa1648e8fde176887bdfedafb075f5bb3e.
//
// Solidity: event Unbond(address indexed validator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) WatchUnbond(opts *bind.WatchOpts, sink chan<- *StakingGeneratedModuleUnbond, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingGeneratedModule.contract.WatchLogs(opts, "Unbond", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingGeneratedModuleUnbond)
				if err := _StakingGeneratedModule.contract.UnpackLog(event, "Unbond", log); err != nil {
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

// ParseUnbond is a log parse operation binding the contract event 0x9b3dc86ff4188cb66e4fbacecb81f0fa1648e8fde176887bdfedafb075f5bb3e.
//
// Solidity: event Unbond(address indexed validator, (uint256,string)[] amount)
func (_StakingGeneratedModule *StakingGeneratedModuleFilterer) ParseUnbond(log types.Log) (*StakingGeneratedModuleUnbond, error) {
	event := new(StakingGeneratedModuleUnbond)
	if err := _StakingGeneratedModule.contract.UnpackLog(event, "Unbond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
