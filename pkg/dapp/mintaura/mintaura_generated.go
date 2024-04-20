// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mintaura

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

// IClaimConditionClaimCondition is an auto generated low-level Go binding around an user-defined struct.
type IClaimConditionClaimCondition struct {
	StartTimestamp         *big.Int
	MaxClaimableSupply     *big.Int
	SupplyClaimed          *big.Int
	QuantityLimitPerWallet *big.Int
	MerkleRoot             [32]byte
	PricePerToken          *big.Int
	Currency               common.Address
	Metadata               string
}

// IDropAllowlistProof is an auto generated low-level Go binding around an user-defined struct.
type IDropAllowlistProof struct {
	Proof                  [][32]byte
	QuantityLimitPerWallet *big.Int
	PricePerToken          *big.Int
	Currency               common.Address
}

// IERC721AUpgradeableTokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type IERC721AUpgradeableTokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
	Burned         bool
	ExtraData      *big.Int
}

// ISharedMetadataSharedMetadataInfo is an auto generated low-level Go binding around an user-defined struct.
type ISharedMetadataSharedMetadataInfo struct {
	Name         string
	Description  string
	ImageURI     string
	AnimationURI string
}

// MintAuraGeneratedMetaData contains all meta data concerning the MintAuraGenerated contract.
var MintAuraGeneratedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQueryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIClaimCondition.ClaimCondition[]\",\"name\":\"claimConditions\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"resetEligibility\",\"type\":\"bool\"}],\"name\":\"ClaimConditionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"prevURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRoyaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoyaltyBps\",\"type\":\"uint256\"}],\"name\":\"DefaultRoyalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"PrimarySaleRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"animationURI\",\"type\":\"string\"}],\"name\":\"SharedMetadataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"claimConditionIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityClaimed\",\"type\":\"uint256\"}],\"name\":\"TokensClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"internalType\":\"structIDrop.AllowlistProof\",\"name\":\"_allowlistProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimCondition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentStartId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"explicitOwnershipOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721AUpgradeable.TokenOwnership\",\"name\":\"ownership\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveClaimConditionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"}],\"name\":\"getClaimConditionById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structIClaimCondition.ClaimCondition\",\"name\":\"condition\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultRoyaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyaltyInfoForToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"}],\"name\":\"getSupplyClaimedByWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyClaimedByWallet\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRoleWithSwitch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_contractURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_trustedForwarders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_saleRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_royaltyBps\",\"type\":\"uint128\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenIdToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenIdToMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"primarySaleRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structIClaimCondition.ClaimCondition[]\",\"name\":\"_conditions\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"_resetClaimEligibility\",\"type\":\"bool\"}],\"name\":\"setClaimConditions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBps\",\"type\":\"uint256\"}],\"name\":\"setDefaultRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_saleRecipient\",\"type\":\"address\"}],\"name\":\"setPrimarySaleRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bps\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyInfoForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animationURI\",\"type\":\"string\"}],\"internalType\":\"structISharedMetadata.SharedMetadataInfo\",\"name\":\"_metadata\",\"type\":\"tuple\"}],\"name\":\"setSharedMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animationURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stop\",\"type\":\"uint256\"}],\"name\":\"tokensOfOwnerIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"internalType\":\"structIDrop.AllowlistProof\",\"name\":\"_allowlistProof\",\"type\":\"tuple\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOverride\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MintAuraGeneratedABI is the input ABI used to generate the binding from.
// Deprecated: Use MintAuraGeneratedMetaData.ABI instead.
var MintAuraGeneratedABI = MintAuraGeneratedMetaData.ABI

// MintAuraGenerated is an auto generated Go binding around an Ethereum contract.
type MintAuraGenerated struct {
	MintAuraGeneratedCaller     // Read-only binding to the contract
	MintAuraGeneratedTransactor // Write-only binding to the contract
	MintAuraGeneratedFilterer   // Log filterer for contract events
}

// MintAuraGeneratedCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintAuraGeneratedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintAuraGeneratedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintAuraGeneratedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintAuraGeneratedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintAuraGeneratedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintAuraGeneratedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintAuraGeneratedSession struct {
	Contract     *MintAuraGenerated // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MintAuraGeneratedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintAuraGeneratedCallerSession struct {
	Contract *MintAuraGeneratedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MintAuraGeneratedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintAuraGeneratedTransactorSession struct {
	Contract     *MintAuraGeneratedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MintAuraGeneratedRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintAuraGeneratedRaw struct {
	Contract *MintAuraGenerated // Generic contract binding to access the raw methods on
}

// MintAuraGeneratedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintAuraGeneratedCallerRaw struct {
	Contract *MintAuraGeneratedCaller // Generic read-only contract binding to access the raw methods on
}

// MintAuraGeneratedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintAuraGeneratedTransactorRaw struct {
	Contract *MintAuraGeneratedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintAuraGenerated creates a new instance of MintAuraGenerated, bound to a specific deployed contract.
func NewMintAuraGenerated(address common.Address, backend bind.ContractBackend) (*MintAuraGenerated, error) {
	contract, err := bindMintAuraGenerated(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintAuraGenerated{MintAuraGeneratedCaller: MintAuraGeneratedCaller{contract: contract}, MintAuraGeneratedTransactor: MintAuraGeneratedTransactor{contract: contract}, MintAuraGeneratedFilterer: MintAuraGeneratedFilterer{contract: contract}}, nil
}

// NewMintAuraGeneratedCaller creates a new read-only instance of MintAuraGenerated, bound to a specific deployed contract.
func NewMintAuraGeneratedCaller(address common.Address, caller bind.ContractCaller) (*MintAuraGeneratedCaller, error) {
	contract, err := bindMintAuraGenerated(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedCaller{contract: contract}, nil
}

// NewMintAuraGeneratedTransactor creates a new write-only instance of MintAuraGenerated, bound to a specific deployed contract.
func NewMintAuraGeneratedTransactor(address common.Address, transactor bind.ContractTransactor) (*MintAuraGeneratedTransactor, error) {
	contract, err := bindMintAuraGenerated(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedTransactor{contract: contract}, nil
}

// NewMintAuraGeneratedFilterer creates a new log filterer instance of MintAuraGenerated, bound to a specific deployed contract.
func NewMintAuraGeneratedFilterer(address common.Address, filterer bind.ContractFilterer) (*MintAuraGeneratedFilterer, error) {
	contract, err := bindMintAuraGenerated(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedFilterer{contract: contract}, nil
}

// bindMintAuraGenerated binds a generic wrapper to an already deployed contract.
func bindMintAuraGenerated(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MintAuraGeneratedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintAuraGenerated *MintAuraGeneratedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintAuraGenerated.Contract.MintAuraGeneratedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintAuraGenerated *MintAuraGeneratedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.MintAuraGeneratedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintAuraGenerated *MintAuraGeneratedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.MintAuraGeneratedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintAuraGenerated *MintAuraGeneratedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintAuraGenerated.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintAuraGenerated *MintAuraGeneratedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintAuraGenerated *MintAuraGeneratedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MintAuraGenerated *MintAuraGeneratedCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MintAuraGenerated *MintAuraGeneratedSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MintAuraGenerated.Contract.DEFAULTADMINROLE(&_MintAuraGenerated.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MintAuraGenerated.Contract.DEFAULTADMINROLE(&_MintAuraGenerated.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MintAuraGenerated.Contract.BalanceOf(&_MintAuraGenerated.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MintAuraGenerated.Contract.BalanceOf(&_MintAuraGenerated.CallOpts, owner)
}

// ClaimCondition is a free data retrieval call binding the contract method 0xd637ed59.
//
// Solidity: function claimCondition() view returns(uint256 currentStartId, uint256 count)
func (_MintAuraGenerated *MintAuraGeneratedCaller) ClaimCondition(opts *bind.CallOpts) (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "claimCondition")

	outstruct := new(struct {
		CurrentStartId *big.Int
		Count          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentStartId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Count = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ClaimCondition is a free data retrieval call binding the contract method 0xd637ed59.
//
// Solidity: function claimCondition() view returns(uint256 currentStartId, uint256 count)
func (_MintAuraGenerated *MintAuraGeneratedSession) ClaimCondition() (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	return _MintAuraGenerated.Contract.ClaimCondition(&_MintAuraGenerated.CallOpts)
}

// ClaimCondition is a free data retrieval call binding the contract method 0xd637ed59.
//
// Solidity: function claimCondition() view returns(uint256 currentStartId, uint256 count)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) ClaimCondition() (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	return _MintAuraGenerated.Contract.ClaimCondition(&_MintAuraGenerated.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedSession) ContractURI() (string, error) {
	return _MintAuraGenerated.Contract.ContractURI(&_MintAuraGenerated.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) ContractURI() (string, error) {
	return _MintAuraGenerated.Contract.ContractURI(&_MintAuraGenerated.CallOpts)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_MintAuraGenerated *MintAuraGeneratedCaller) ExplicitOwnershipOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "explicitOwnershipOf", tokenId)

	if err != nil {
		return *new(IERC721AUpgradeableTokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721AUpgradeableTokenOwnership)).(*IERC721AUpgradeableTokenOwnership)

	return out0, err

}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_MintAuraGenerated *MintAuraGeneratedSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	return _MintAuraGenerated.Contract.ExplicitOwnershipOf(&_MintAuraGenerated.CallOpts, tokenId)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	return _MintAuraGenerated.Contract.ExplicitOwnershipOf(&_MintAuraGenerated.CallOpts, tokenId)
}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0xc68907de.
//
// Solidity: function getActiveClaimConditionId() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCaller) GetActiveClaimConditionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "getActiveClaimConditionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0xc68907de.
//
// Solidity: function getActiveClaimConditionId() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedSession) GetActiveClaimConditionId() (*big.Int, error) {
	return _MintAuraGenerated.Contract.GetActiveClaimConditionId(&_MintAuraGenerated.CallOpts)
}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0xc68907de.
//
// Solidity: function getActiveClaimConditionId() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) GetActiveClaimConditionId() (*big.Int, error) {
	return _MintAuraGenerated.Contract.GetActiveClaimConditionId(&_MintAuraGenerated.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MintAuraGenerated.Contract.GetApproved(&_MintAuraGenerated.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MintAuraGenerated.Contract.GetApproved(&_MintAuraGenerated.CallOpts, tokenId)
}

// GetClaimConditionById is a free data retrieval call binding the contract method 0x6f8934f4.
//
// Solidity: function getClaimConditionById(uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,bytes32,uint256,address,string) condition)
func (_MintAuraGenerated *MintAuraGeneratedCaller) GetClaimConditionById(opts *bind.CallOpts, _conditionId *big.Int) (IClaimConditionClaimCondition, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "getClaimConditionById", _conditionId)

	if err != nil {
		return *new(IClaimConditionClaimCondition), err
	}

	out0 := *abi.ConvertType(out[0], new(IClaimConditionClaimCondition)).(*IClaimConditionClaimCondition)

	return out0, err

}

// GetClaimConditionById is a free data retrieval call binding the contract method 0x6f8934f4.
//
// Solidity: function getClaimConditionById(uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,bytes32,uint256,address,string) condition)
func (_MintAuraGenerated *MintAuraGeneratedSession) GetClaimConditionById(_conditionId *big.Int) (IClaimConditionClaimCondition, error) {
	return _MintAuraGenerated.Contract.GetClaimConditionById(&_MintAuraGenerated.CallOpts, _conditionId)
}

// GetClaimConditionById is a free data retrieval call binding the contract method 0x6f8934f4.
//
// Solidity: function getClaimConditionById(uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,bytes32,uint256,address,string) condition)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) GetClaimConditionById(_conditionId *big.Int) (IClaimConditionClaimCondition, error) {
	return _MintAuraGenerated.Contract.GetClaimConditionById(&_MintAuraGenerated.CallOpts, _conditionId)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_MintAuraGenerated *MintAuraGeneratedCaller) GetDefaultRoyaltyInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "getDefaultRoyaltyInfo")

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_MintAuraGenerated *MintAuraGeneratedSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _MintAuraGenerated.Contract.GetDefaultRoyaltyInfo(&_MintAuraGenerated.CallOpts)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _MintAuraGenerated.Contract.GetDefaultRoyaltyInfo(&_MintAuraGenerated.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MintAuraGenerated *MintAuraGeneratedCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MintAuraGenerated *MintAuraGeneratedSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MintAuraGenerated.Contract.GetRoleAdmin(&_MintAuraGenerated.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MintAuraGenerated.Contract.GetRoleAdmin(&_MintAuraGenerated.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_MintAuraGenerated *MintAuraGeneratedCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_MintAuraGenerated *MintAuraGeneratedSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MintAuraGenerated.Contract.GetRoleMember(&_MintAuraGenerated.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MintAuraGenerated.Contract.GetRoleMember(&_MintAuraGenerated.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_MintAuraGenerated *MintAuraGeneratedCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_MintAuraGenerated *MintAuraGeneratedSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MintAuraGenerated.Contract.GetRoleMemberCount(&_MintAuraGenerated.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MintAuraGenerated.Contract.GetRoleMemberCount(&_MintAuraGenerated.CallOpts, role)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_MintAuraGenerated *MintAuraGeneratedCaller) GetRoyaltyInfoForToken(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, uint16, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "getRoyaltyInfoForToken", _tokenId)

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_MintAuraGenerated *MintAuraGeneratedSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _MintAuraGenerated.Contract.GetRoyaltyInfoForToken(&_MintAuraGenerated.CallOpts, _tokenId)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _MintAuraGenerated.Contract.GetRoyaltyInfoForToken(&_MintAuraGenerated.CallOpts, _tokenId)
}

// GetSupplyClaimedByWallet is a free data retrieval call binding the contract method 0xad1eefc5.
//
// Solidity: function getSupplyClaimedByWallet(uint256 _conditionId, address _claimer) view returns(uint256 supplyClaimedByWallet)
func (_MintAuraGenerated *MintAuraGeneratedCaller) GetSupplyClaimedByWallet(opts *bind.CallOpts, _conditionId *big.Int, _claimer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "getSupplyClaimedByWallet", _conditionId, _claimer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSupplyClaimedByWallet is a free data retrieval call binding the contract method 0xad1eefc5.
//
// Solidity: function getSupplyClaimedByWallet(uint256 _conditionId, address _claimer) view returns(uint256 supplyClaimedByWallet)
func (_MintAuraGenerated *MintAuraGeneratedSession) GetSupplyClaimedByWallet(_conditionId *big.Int, _claimer common.Address) (*big.Int, error) {
	return _MintAuraGenerated.Contract.GetSupplyClaimedByWallet(&_MintAuraGenerated.CallOpts, _conditionId, _claimer)
}

// GetSupplyClaimedByWallet is a free data retrieval call binding the contract method 0xad1eefc5.
//
// Solidity: function getSupplyClaimedByWallet(uint256 _conditionId, address _claimer) view returns(uint256 supplyClaimedByWallet)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) GetSupplyClaimedByWallet(_conditionId *big.Int, _claimer common.Address) (*big.Int, error) {
	return _MintAuraGenerated.Contract.GetSupplyClaimedByWallet(&_MintAuraGenerated.CallOpts, _conditionId, _claimer)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MintAuraGenerated.Contract.HasRole(&_MintAuraGenerated.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MintAuraGenerated.Contract.HasRole(&_MintAuraGenerated.CallOpts, role, account)
}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCaller) HasRoleWithSwitch(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "hasRoleWithSwitch", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedSession) HasRoleWithSwitch(role [32]byte, account common.Address) (bool, error) {
	return _MintAuraGenerated.Contract.HasRoleWithSwitch(&_MintAuraGenerated.CallOpts, role, account)
}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) HasRoleWithSwitch(role [32]byte, account common.Address) (bool, error) {
	return _MintAuraGenerated.Contract.HasRoleWithSwitch(&_MintAuraGenerated.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MintAuraGenerated.Contract.IsApprovedForAll(&_MintAuraGenerated.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MintAuraGenerated.Contract.IsApprovedForAll(&_MintAuraGenerated.CallOpts, owner, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _MintAuraGenerated.Contract.IsTrustedForwarder(&_MintAuraGenerated.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _MintAuraGenerated.Contract.IsTrustedForwarder(&_MintAuraGenerated.CallOpts, forwarder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedSession) Name() (string, error) {
	return _MintAuraGenerated.Contract.Name(&_MintAuraGenerated.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) Name() (string, error) {
	return _MintAuraGenerated.Contract.Name(&_MintAuraGenerated.CallOpts)
}

// NextTokenIdToClaim is a free data retrieval call binding the contract method 0xacd083f8.
//
// Solidity: function nextTokenIdToClaim() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCaller) NextTokenIdToClaim(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "nextTokenIdToClaim")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToClaim is a free data retrieval call binding the contract method 0xacd083f8.
//
// Solidity: function nextTokenIdToClaim() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedSession) NextTokenIdToClaim() (*big.Int, error) {
	return _MintAuraGenerated.Contract.NextTokenIdToClaim(&_MintAuraGenerated.CallOpts)
}

// NextTokenIdToClaim is a free data retrieval call binding the contract method 0xacd083f8.
//
// Solidity: function nextTokenIdToClaim() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) NextTokenIdToClaim() (*big.Int, error) {
	return _MintAuraGenerated.Contract.NextTokenIdToClaim(&_MintAuraGenerated.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCaller) NextTokenIdToMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "nextTokenIdToMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedSession) NextTokenIdToMint() (*big.Int, error) {
	return _MintAuraGenerated.Contract.NextTokenIdToMint(&_MintAuraGenerated.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) NextTokenIdToMint() (*big.Int, error) {
	return _MintAuraGenerated.Contract.NextTokenIdToMint(&_MintAuraGenerated.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedSession) Owner() (common.Address, error) {
	return _MintAuraGenerated.Contract.Owner(&_MintAuraGenerated.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) Owner() (common.Address, error) {
	return _MintAuraGenerated.Contract.Owner(&_MintAuraGenerated.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MintAuraGenerated.Contract.OwnerOf(&_MintAuraGenerated.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MintAuraGenerated.Contract.OwnerOf(&_MintAuraGenerated.CallOpts, tokenId)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedCaller) PrimarySaleRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "primarySaleRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedSession) PrimarySaleRecipient() (common.Address, error) {
	return _MintAuraGenerated.Contract.PrimarySaleRecipient(&_MintAuraGenerated.CallOpts)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) PrimarySaleRecipient() (common.Address, error) {
	return _MintAuraGenerated.Contract.PrimarySaleRecipient(&_MintAuraGenerated.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_MintAuraGenerated *MintAuraGeneratedCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_MintAuraGenerated *MintAuraGeneratedSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _MintAuraGenerated.Contract.RoyaltyInfo(&_MintAuraGenerated.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _MintAuraGenerated.Contract.RoyaltyInfo(&_MintAuraGenerated.CallOpts, tokenId, salePrice)
}

// SharedMetadata is a free data retrieval call binding the contract method 0xb280f703.
//
// Solidity: function sharedMetadata() view returns(string name, string description, string imageURI, string animationURI)
func (_MintAuraGenerated *MintAuraGeneratedCaller) SharedMetadata(opts *bind.CallOpts) (struct {
	Name         string
	Description  string
	ImageURI     string
	AnimationURI string
}, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "sharedMetadata")

	outstruct := new(struct {
		Name         string
		Description  string
		ImageURI     string
		AnimationURI string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ImageURI = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.AnimationURI = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// SharedMetadata is a free data retrieval call binding the contract method 0xb280f703.
//
// Solidity: function sharedMetadata() view returns(string name, string description, string imageURI, string animationURI)
func (_MintAuraGenerated *MintAuraGeneratedSession) SharedMetadata() (struct {
	Name         string
	Description  string
	ImageURI     string
	AnimationURI string
}, error) {
	return _MintAuraGenerated.Contract.SharedMetadata(&_MintAuraGenerated.CallOpts)
}

// SharedMetadata is a free data retrieval call binding the contract method 0xb280f703.
//
// Solidity: function sharedMetadata() view returns(string name, string description, string imageURI, string animationURI)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) SharedMetadata() (struct {
	Name         string
	Description  string
	ImageURI     string
	AnimationURI string
}, error) {
	return _MintAuraGenerated.Contract.SharedMetadata(&_MintAuraGenerated.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() pure returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCaller) StartTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "startTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() pure returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedSession) StartTokenId() (*big.Int, error) {
	return _MintAuraGenerated.Contract.StartTokenId(&_MintAuraGenerated.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() pure returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) StartTokenId() (*big.Int, error) {
	return _MintAuraGenerated.Contract.StartTokenId(&_MintAuraGenerated.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MintAuraGenerated.Contract.SupportsInterface(&_MintAuraGenerated.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MintAuraGenerated.Contract.SupportsInterface(&_MintAuraGenerated.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedSession) Symbol() (string, error) {
	return _MintAuraGenerated.Contract.Symbol(&_MintAuraGenerated.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) Symbol() (string, error) {
	return _MintAuraGenerated.Contract.Symbol(&_MintAuraGenerated.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _MintAuraGenerated.Contract.TokenURI(&_MintAuraGenerated.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _MintAuraGenerated.Contract.TokenURI(&_MintAuraGenerated.CallOpts, _tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_MintAuraGenerated *MintAuraGeneratedCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_MintAuraGenerated *MintAuraGeneratedSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _MintAuraGenerated.Contract.TokensOfOwner(&_MintAuraGenerated.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _MintAuraGenerated.Contract.TokensOfOwner(&_MintAuraGenerated.CallOpts, owner)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_MintAuraGenerated *MintAuraGeneratedCaller) TokensOfOwnerIn(opts *bind.CallOpts, owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "tokensOfOwnerIn", owner, start, stop)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_MintAuraGenerated *MintAuraGeneratedSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _MintAuraGenerated.Contract.TokensOfOwnerIn(&_MintAuraGenerated.CallOpts, owner, start, stop)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _MintAuraGenerated.Contract.TokensOfOwnerIn(&_MintAuraGenerated.CallOpts, owner, start, stop)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedSession) TotalMinted() (*big.Int, error) {
	return _MintAuraGenerated.Contract.TotalMinted(&_MintAuraGenerated.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) TotalMinted() (*big.Int, error) {
	return _MintAuraGenerated.Contract.TotalMinted(&_MintAuraGenerated.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedSession) TotalSupply() (*big.Int, error) {
	return _MintAuraGenerated.Contract.TotalSupply(&_MintAuraGenerated.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) TotalSupply() (*big.Int, error) {
	return _MintAuraGenerated.Contract.TotalSupply(&_MintAuraGenerated.CallOpts)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x23a2902b.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof) view returns(bool isOverride)
func (_MintAuraGenerated *MintAuraGeneratedCaller) VerifyClaim(opts *bind.CallOpts, _conditionId *big.Int, _claimer common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof) (bool, error) {
	var out []interface{}
	err := _MintAuraGenerated.contract.Call(opts, &out, "verifyClaim", _conditionId, _claimer, _quantity, _currency, _pricePerToken, _allowlistProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClaim is a free data retrieval call binding the contract method 0x23a2902b.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof) view returns(bool isOverride)
func (_MintAuraGenerated *MintAuraGeneratedSession) VerifyClaim(_conditionId *big.Int, _claimer common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof) (bool, error) {
	return _MintAuraGenerated.Contract.VerifyClaim(&_MintAuraGenerated.CallOpts, _conditionId, _claimer, _quantity, _currency, _pricePerToken, _allowlistProof)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x23a2902b.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof) view returns(bool isOverride)
func (_MintAuraGenerated *MintAuraGeneratedCallerSession) VerifyClaim(_conditionId *big.Int, _claimer common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof) (bool, error) {
	return _MintAuraGenerated.Contract.VerifyClaim(&_MintAuraGenerated.CallOpts, _conditionId, _claimer, _quantity, _currency, _pricePerToken, _allowlistProof)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Approve(&_MintAuraGenerated.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Approve(&_MintAuraGenerated.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Burn(&_MintAuraGenerated.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Burn(&_MintAuraGenerated.TransactOpts, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x84bb1e42.
//
// Solidity: function claim(address _receiver, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof, bytes _data) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof, _data []byte) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "claim", _receiver, _quantity, _currency, _pricePerToken, _allowlistProof, _data)
}

// Claim is a paid mutator transaction binding the contract method 0x84bb1e42.
//
// Solidity: function claim(address _receiver, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof, bytes _data) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) Claim(_receiver common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof, _data []byte) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Claim(&_MintAuraGenerated.TransactOpts, _receiver, _quantity, _currency, _pricePerToken, _allowlistProof, _data)
}

// Claim is a paid mutator transaction binding the contract method 0x84bb1e42.
//
// Solidity: function claim(address _receiver, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof, bytes _data) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) Claim(_receiver common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof, _data []byte) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Claim(&_MintAuraGenerated.TransactOpts, _receiver, _quantity, _currency, _pricePerToken, _allowlistProof, _data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.GrantRole(&_MintAuraGenerated.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.GrantRole(&_MintAuraGenerated.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x49c5c5b6.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "initialize", _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps)
}

// Initialize is a paid mutator transaction binding the contract method 0x49c5c5b6.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Initialize(&_MintAuraGenerated.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps)
}

// Initialize is a paid mutator transaction binding the contract method 0x49c5c5b6.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Initialize(&_MintAuraGenerated.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MintAuraGenerated *MintAuraGeneratedTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MintAuraGenerated *MintAuraGeneratedSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Multicall(&_MintAuraGenerated.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.Multicall(&_MintAuraGenerated.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.RenounceRole(&_MintAuraGenerated.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.RenounceRole(&_MintAuraGenerated.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.RevokeRole(&_MintAuraGenerated.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.RevokeRole(&_MintAuraGenerated.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SafeTransferFrom(&_MintAuraGenerated.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SafeTransferFrom(&_MintAuraGenerated.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SafeTransferFrom0(&_MintAuraGenerated.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SafeTransferFrom0(&_MintAuraGenerated.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetApprovalForAll(&_MintAuraGenerated.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetApprovalForAll(&_MintAuraGenerated.TransactOpts, operator, approved)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0x74bc7db7.
//
// Solidity: function setClaimConditions((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] _conditions, bool _resetClaimEligibility) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SetClaimConditions(opts *bind.TransactOpts, _conditions []IClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "setClaimConditions", _conditions, _resetClaimEligibility)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0x74bc7db7.
//
// Solidity: function setClaimConditions((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] _conditions, bool _resetClaimEligibility) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SetClaimConditions(_conditions []IClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetClaimConditions(&_MintAuraGenerated.TransactOpts, _conditions, _resetClaimEligibility)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0x74bc7db7.
//
// Solidity: function setClaimConditions((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] _conditions, bool _resetClaimEligibility) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SetClaimConditions(_conditions []IClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetClaimConditions(&_MintAuraGenerated.TransactOpts, _conditions, _resetClaimEligibility)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetContractURI(&_MintAuraGenerated.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetContractURI(&_MintAuraGenerated.TransactOpts, _uri)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SetDefaultRoyaltyInfo(opts *bind.TransactOpts, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "setDefaultRoyaltyInfo", _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetDefaultRoyaltyInfo(&_MintAuraGenerated.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetDefaultRoyaltyInfo(&_MintAuraGenerated.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetOwner(&_MintAuraGenerated.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetOwner(&_MintAuraGenerated.TransactOpts, _newOwner)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SetPrimarySaleRecipient(opts *bind.TransactOpts, _saleRecipient common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "setPrimarySaleRecipient", _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetPrimarySaleRecipient(&_MintAuraGenerated.TransactOpts, _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetPrimarySaleRecipient(&_MintAuraGenerated.TransactOpts, _saleRecipient)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SetRoyaltyInfoForToken(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "setRoyaltyInfoForToken", _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetRoyaltyInfoForToken(&_MintAuraGenerated.TransactOpts, _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetRoyaltyInfoForToken(&_MintAuraGenerated.TransactOpts, _tokenId, _recipient, _bps)
}

// SetSharedMetadata is a paid mutator transaction binding the contract method 0xa7d27d9d.
//
// Solidity: function setSharedMetadata((string,string,string,string) _metadata) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) SetSharedMetadata(opts *bind.TransactOpts, _metadata ISharedMetadataSharedMetadataInfo) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "setSharedMetadata", _metadata)
}

// SetSharedMetadata is a paid mutator transaction binding the contract method 0xa7d27d9d.
//
// Solidity: function setSharedMetadata((string,string,string,string) _metadata) returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) SetSharedMetadata(_metadata ISharedMetadataSharedMetadataInfo) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetSharedMetadata(&_MintAuraGenerated.TransactOpts, _metadata)
}

// SetSharedMetadata is a paid mutator transaction binding the contract method 0xa7d27d9d.
//
// Solidity: function setSharedMetadata((string,string,string,string) _metadata) returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) SetSharedMetadata(_metadata ISharedMetadataSharedMetadataInfo) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.SetSharedMetadata(&_MintAuraGenerated.TransactOpts, _metadata)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.TransferFrom(&_MintAuraGenerated.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_MintAuraGenerated *MintAuraGeneratedTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MintAuraGenerated.Contract.TransferFrom(&_MintAuraGenerated.TransactOpts, from, to, tokenId)
}

// MintAuraGeneratedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MintAuraGenerated contract.
type MintAuraGeneratedApprovalIterator struct {
	Event *MintAuraGeneratedApproval // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedApproval)
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
		it.Event = new(MintAuraGeneratedApproval)
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
func (it *MintAuraGeneratedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedApproval represents a Approval event raised by the MintAuraGenerated contract.
type MintAuraGeneratedApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MintAuraGeneratedApprovalIterator, error) {

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

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedApprovalIterator{contract: _MintAuraGenerated.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedApproval)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseApproval(log types.Log) (*MintAuraGeneratedApproval, error) {
	event := new(MintAuraGeneratedApproval)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MintAuraGenerated contract.
type MintAuraGeneratedApprovalForAllIterator struct {
	Event *MintAuraGeneratedApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedApprovalForAll)
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
		it.Event = new(MintAuraGeneratedApprovalForAll)
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
func (it *MintAuraGeneratedApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedApprovalForAll represents a ApprovalForAll event raised by the MintAuraGenerated contract.
type MintAuraGeneratedApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MintAuraGeneratedApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedApprovalForAllIterator{contract: _MintAuraGenerated.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedApprovalForAll)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseApprovalForAll(log types.Log) (*MintAuraGeneratedApprovalForAll, error) {
	event := new(MintAuraGeneratedApprovalForAll)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the MintAuraGenerated contract.
type MintAuraGeneratedBatchMetadataUpdateIterator struct {
	Event *MintAuraGeneratedBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedBatchMetadataUpdate)
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
		it.Event = new(MintAuraGeneratedBatchMetadataUpdate)
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
func (it *MintAuraGeneratedBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the MintAuraGenerated contract.
type MintAuraGeneratedBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*MintAuraGeneratedBatchMetadataUpdateIterator, error) {

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedBatchMetadataUpdateIterator{contract: _MintAuraGenerated.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedBatchMetadataUpdate)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseBatchMetadataUpdate(log types.Log) (*MintAuraGeneratedBatchMetadataUpdate, error) {
	event := new(MintAuraGeneratedBatchMetadataUpdate)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedClaimConditionsUpdatedIterator is returned from FilterClaimConditionsUpdated and is used to iterate over the raw logs and unpacked data for ClaimConditionsUpdated events raised by the MintAuraGenerated contract.
type MintAuraGeneratedClaimConditionsUpdatedIterator struct {
	Event *MintAuraGeneratedClaimConditionsUpdated // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedClaimConditionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedClaimConditionsUpdated)
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
		it.Event = new(MintAuraGeneratedClaimConditionsUpdated)
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
func (it *MintAuraGeneratedClaimConditionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedClaimConditionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedClaimConditionsUpdated represents a ClaimConditionsUpdated event raised by the MintAuraGenerated contract.
type MintAuraGeneratedClaimConditionsUpdated struct {
	ClaimConditions  []IClaimConditionClaimCondition
	ResetEligibility bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimConditionsUpdated is a free log retrieval operation binding the contract event 0xbf4016fceeaaa4ac5cf4be865b559ff85825ab4ca7aa7b661d16e2f544c03098.
//
// Solidity: event ClaimConditionsUpdated((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] claimConditions, bool resetEligibility)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterClaimConditionsUpdated(opts *bind.FilterOpts) (*MintAuraGeneratedClaimConditionsUpdatedIterator, error) {

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "ClaimConditionsUpdated")
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedClaimConditionsUpdatedIterator{contract: _MintAuraGenerated.contract, event: "ClaimConditionsUpdated", logs: logs, sub: sub}, nil
}

// WatchClaimConditionsUpdated is a free log subscription operation binding the contract event 0xbf4016fceeaaa4ac5cf4be865b559ff85825ab4ca7aa7b661d16e2f544c03098.
//
// Solidity: event ClaimConditionsUpdated((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] claimConditions, bool resetEligibility)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchClaimConditionsUpdated(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedClaimConditionsUpdated) (event.Subscription, error) {

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "ClaimConditionsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedClaimConditionsUpdated)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "ClaimConditionsUpdated", log); err != nil {
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

// ParseClaimConditionsUpdated is a log parse operation binding the contract event 0xbf4016fceeaaa4ac5cf4be865b559ff85825ab4ca7aa7b661d16e2f544c03098.
//
// Solidity: event ClaimConditionsUpdated((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] claimConditions, bool resetEligibility)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseClaimConditionsUpdated(log types.Log) (*MintAuraGeneratedClaimConditionsUpdated, error) {
	event := new(MintAuraGeneratedClaimConditionsUpdated)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "ClaimConditionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the MintAuraGenerated contract.
type MintAuraGeneratedConsecutiveTransferIterator struct {
	Event *MintAuraGeneratedConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedConsecutiveTransfer)
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
		it.Event = new(MintAuraGeneratedConsecutiveTransfer)
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
func (it *MintAuraGeneratedConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedConsecutiveTransfer represents a ConsecutiveTransfer event raised by the MintAuraGenerated contract.
type MintAuraGeneratedConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*MintAuraGeneratedConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedConsecutiveTransferIterator{contract: _MintAuraGenerated.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedConsecutiveTransfer)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseConsecutiveTransfer(log types.Log) (*MintAuraGeneratedConsecutiveTransfer, error) {
	event := new(MintAuraGeneratedConsecutiveTransfer)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the MintAuraGenerated contract.
type MintAuraGeneratedContractURIUpdatedIterator struct {
	Event *MintAuraGeneratedContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedContractURIUpdated)
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
		it.Event = new(MintAuraGeneratedContractURIUpdated)
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
func (it *MintAuraGeneratedContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedContractURIUpdated represents a ContractURIUpdated event raised by the MintAuraGenerated contract.
type MintAuraGeneratedContractURIUpdated struct {
	PrevURI string
	NewURI  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*MintAuraGeneratedContractURIUpdatedIterator, error) {

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedContractURIUpdatedIterator{contract: _MintAuraGenerated.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedContractURIUpdated)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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

// ParseContractURIUpdated is a log parse operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseContractURIUpdated(log types.Log) (*MintAuraGeneratedContractURIUpdated, error) {
	event := new(MintAuraGeneratedContractURIUpdated)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedDefaultRoyaltyIterator is returned from FilterDefaultRoyalty and is used to iterate over the raw logs and unpacked data for DefaultRoyalty events raised by the MintAuraGenerated contract.
type MintAuraGeneratedDefaultRoyaltyIterator struct {
	Event *MintAuraGeneratedDefaultRoyalty // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedDefaultRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedDefaultRoyalty)
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
		it.Event = new(MintAuraGeneratedDefaultRoyalty)
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
func (it *MintAuraGeneratedDefaultRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedDefaultRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedDefaultRoyalty represents a DefaultRoyalty event raised by the MintAuraGenerated contract.
type MintAuraGeneratedDefaultRoyalty struct {
	NewRoyaltyRecipient common.Address
	NewRoyaltyBps       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyalty is a free log retrieval operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterDefaultRoyalty(opts *bind.FilterOpts, newRoyaltyRecipient []common.Address) (*MintAuraGeneratedDefaultRoyaltyIterator, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedDefaultRoyaltyIterator{contract: _MintAuraGenerated.contract, event: "DefaultRoyalty", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyalty is a free log subscription operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchDefaultRoyalty(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedDefaultRoyalty, newRoyaltyRecipient []common.Address) (event.Subscription, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedDefaultRoyalty)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
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

// ParseDefaultRoyalty is a log parse operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseDefaultRoyalty(log types.Log) (*MintAuraGeneratedDefaultRoyalty, error) {
	event := new(MintAuraGeneratedDefaultRoyalty)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MintAuraGenerated contract.
type MintAuraGeneratedInitializedIterator struct {
	Event *MintAuraGeneratedInitialized // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedInitialized)
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
		it.Event = new(MintAuraGeneratedInitialized)
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
func (it *MintAuraGeneratedInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedInitialized represents a Initialized event raised by the MintAuraGenerated contract.
type MintAuraGeneratedInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterInitialized(opts *bind.FilterOpts) (*MintAuraGeneratedInitializedIterator, error) {

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedInitializedIterator{contract: _MintAuraGenerated.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedInitialized) (event.Subscription, error) {

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedInitialized)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseInitialized(log types.Log) (*MintAuraGeneratedInitialized, error) {
	event := new(MintAuraGeneratedInitialized)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the MintAuraGenerated contract.
type MintAuraGeneratedMetadataUpdateIterator struct {
	Event *MintAuraGeneratedMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedMetadataUpdate)
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
		it.Event = new(MintAuraGeneratedMetadataUpdate)
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
func (it *MintAuraGeneratedMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedMetadataUpdate represents a MetadataUpdate event raised by the MintAuraGenerated contract.
type MintAuraGeneratedMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*MintAuraGeneratedMetadataUpdateIterator, error) {

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedMetadataUpdateIterator{contract: _MintAuraGenerated.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedMetadataUpdate)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseMetadataUpdate(log types.Log) (*MintAuraGeneratedMetadataUpdate, error) {
	event := new(MintAuraGeneratedMetadataUpdate)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedOwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the MintAuraGenerated contract.
type MintAuraGeneratedOwnerUpdatedIterator struct {
	Event *MintAuraGeneratedOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedOwnerUpdated)
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
		it.Event = new(MintAuraGeneratedOwnerUpdated)
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
func (it *MintAuraGeneratedOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedOwnerUpdated represents a OwnerUpdated event raised by the MintAuraGenerated contract.
type MintAuraGeneratedOwnerUpdated struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterOwnerUpdated(opts *bind.FilterOpts, prevOwner []common.Address, newOwner []common.Address) (*MintAuraGeneratedOwnerUpdatedIterator, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedOwnerUpdatedIterator{contract: _MintAuraGenerated.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedOwnerUpdated, prevOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedOwnerUpdated)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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

// ParseOwnerUpdated is a log parse operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseOwnerUpdated(log types.Log) (*MintAuraGeneratedOwnerUpdated, error) {
	event := new(MintAuraGeneratedOwnerUpdated)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedPrimarySaleRecipientUpdatedIterator is returned from FilterPrimarySaleRecipientUpdated and is used to iterate over the raw logs and unpacked data for PrimarySaleRecipientUpdated events raised by the MintAuraGenerated contract.
type MintAuraGeneratedPrimarySaleRecipientUpdatedIterator struct {
	Event *MintAuraGeneratedPrimarySaleRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedPrimarySaleRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedPrimarySaleRecipientUpdated)
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
		it.Event = new(MintAuraGeneratedPrimarySaleRecipientUpdated)
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
func (it *MintAuraGeneratedPrimarySaleRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedPrimarySaleRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedPrimarySaleRecipientUpdated represents a PrimarySaleRecipientUpdated event raised by the MintAuraGenerated contract.
type MintAuraGeneratedPrimarySaleRecipientUpdated struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrimarySaleRecipientUpdated is a free log retrieval operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterPrimarySaleRecipientUpdated(opts *bind.FilterOpts, recipient []common.Address) (*MintAuraGeneratedPrimarySaleRecipientUpdatedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedPrimarySaleRecipientUpdatedIterator{contract: _MintAuraGenerated.contract, event: "PrimarySaleRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchPrimarySaleRecipientUpdated is a free log subscription operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchPrimarySaleRecipientUpdated(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedPrimarySaleRecipientUpdated, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedPrimarySaleRecipientUpdated)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
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

// ParsePrimarySaleRecipientUpdated is a log parse operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParsePrimarySaleRecipientUpdated(log types.Log) (*MintAuraGeneratedPrimarySaleRecipientUpdated, error) {
	event := new(MintAuraGeneratedPrimarySaleRecipientUpdated)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MintAuraGenerated contract.
type MintAuraGeneratedRoleAdminChangedIterator struct {
	Event *MintAuraGeneratedRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedRoleAdminChanged)
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
		it.Event = new(MintAuraGeneratedRoleAdminChanged)
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
func (it *MintAuraGeneratedRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedRoleAdminChanged represents a RoleAdminChanged event raised by the MintAuraGenerated contract.
type MintAuraGeneratedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MintAuraGeneratedRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedRoleAdminChangedIterator{contract: _MintAuraGenerated.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedRoleAdminChanged)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseRoleAdminChanged(log types.Log) (*MintAuraGeneratedRoleAdminChanged, error) {
	event := new(MintAuraGeneratedRoleAdminChanged)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MintAuraGenerated contract.
type MintAuraGeneratedRoleGrantedIterator struct {
	Event *MintAuraGeneratedRoleGranted // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedRoleGranted)
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
		it.Event = new(MintAuraGeneratedRoleGranted)
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
func (it *MintAuraGeneratedRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedRoleGranted represents a RoleGranted event raised by the MintAuraGenerated contract.
type MintAuraGeneratedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MintAuraGeneratedRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedRoleGrantedIterator{contract: _MintAuraGenerated.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedRoleGranted)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseRoleGranted(log types.Log) (*MintAuraGeneratedRoleGranted, error) {
	event := new(MintAuraGeneratedRoleGranted)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MintAuraGenerated contract.
type MintAuraGeneratedRoleRevokedIterator struct {
	Event *MintAuraGeneratedRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedRoleRevoked)
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
		it.Event = new(MintAuraGeneratedRoleRevoked)
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
func (it *MintAuraGeneratedRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedRoleRevoked represents a RoleRevoked event raised by the MintAuraGenerated contract.
type MintAuraGeneratedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MintAuraGeneratedRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedRoleRevokedIterator{contract: _MintAuraGenerated.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedRoleRevoked)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseRoleRevoked(log types.Log) (*MintAuraGeneratedRoleRevoked, error) {
	event := new(MintAuraGeneratedRoleRevoked)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedRoyaltyForTokenIterator is returned from FilterRoyaltyForToken and is used to iterate over the raw logs and unpacked data for RoyaltyForToken events raised by the MintAuraGenerated contract.
type MintAuraGeneratedRoyaltyForTokenIterator struct {
	Event *MintAuraGeneratedRoyaltyForToken // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedRoyaltyForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedRoyaltyForToken)
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
		it.Event = new(MintAuraGeneratedRoyaltyForToken)
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
func (it *MintAuraGeneratedRoyaltyForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedRoyaltyForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedRoyaltyForToken represents a RoyaltyForToken event raised by the MintAuraGenerated contract.
type MintAuraGeneratedRoyaltyForToken struct {
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	RoyaltyBps       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyForToken is a free log retrieval operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterRoyaltyForToken(opts *bind.FilterOpts, tokenId []*big.Int, royaltyRecipient []common.Address) (*MintAuraGeneratedRoyaltyForTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedRoyaltyForTokenIterator{contract: _MintAuraGenerated.contract, event: "RoyaltyForToken", logs: logs, sub: sub}, nil
}

// WatchRoyaltyForToken is a free log subscription operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchRoyaltyForToken(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedRoyaltyForToken, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedRoyaltyForToken)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
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

// ParseRoyaltyForToken is a log parse operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseRoyaltyForToken(log types.Log) (*MintAuraGeneratedRoyaltyForToken, error) {
	event := new(MintAuraGeneratedRoyaltyForToken)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedSharedMetadataUpdatedIterator is returned from FilterSharedMetadataUpdated and is used to iterate over the raw logs and unpacked data for SharedMetadataUpdated events raised by the MintAuraGenerated contract.
type MintAuraGeneratedSharedMetadataUpdatedIterator struct {
	Event *MintAuraGeneratedSharedMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedSharedMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedSharedMetadataUpdated)
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
		it.Event = new(MintAuraGeneratedSharedMetadataUpdated)
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
func (it *MintAuraGeneratedSharedMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedSharedMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedSharedMetadataUpdated represents a SharedMetadataUpdated event raised by the MintAuraGenerated contract.
type MintAuraGeneratedSharedMetadataUpdated struct {
	Name         string
	Description  string
	ImageURI     string
	AnimationURI string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSharedMetadataUpdated is a free log retrieval operation binding the contract event 0x8edd7f36d5f01bd45e59cf55b0a670dcf701fc20f678970a8c243b2346d6acaf.
//
// Solidity: event SharedMetadataUpdated(string name, string description, string imageURI, string animationURI)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterSharedMetadataUpdated(opts *bind.FilterOpts) (*MintAuraGeneratedSharedMetadataUpdatedIterator, error) {

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "SharedMetadataUpdated")
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedSharedMetadataUpdatedIterator{contract: _MintAuraGenerated.contract, event: "SharedMetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchSharedMetadataUpdated is a free log subscription operation binding the contract event 0x8edd7f36d5f01bd45e59cf55b0a670dcf701fc20f678970a8c243b2346d6acaf.
//
// Solidity: event SharedMetadataUpdated(string name, string description, string imageURI, string animationURI)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchSharedMetadataUpdated(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedSharedMetadataUpdated) (event.Subscription, error) {

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "SharedMetadataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedSharedMetadataUpdated)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "SharedMetadataUpdated", log); err != nil {
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

// ParseSharedMetadataUpdated is a log parse operation binding the contract event 0x8edd7f36d5f01bd45e59cf55b0a670dcf701fc20f678970a8c243b2346d6acaf.
//
// Solidity: event SharedMetadataUpdated(string name, string description, string imageURI, string animationURI)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseSharedMetadataUpdated(log types.Log) (*MintAuraGeneratedSharedMetadataUpdated, error) {
	event := new(MintAuraGeneratedSharedMetadataUpdated)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "SharedMetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedTokensClaimedIterator is returned from FilterTokensClaimed and is used to iterate over the raw logs and unpacked data for TokensClaimed events raised by the MintAuraGenerated contract.
type MintAuraGeneratedTokensClaimedIterator struct {
	Event *MintAuraGeneratedTokensClaimed // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedTokensClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedTokensClaimed)
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
		it.Event = new(MintAuraGeneratedTokensClaimed)
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
func (it *MintAuraGeneratedTokensClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedTokensClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedTokensClaimed represents a TokensClaimed event raised by the MintAuraGenerated contract.
type MintAuraGeneratedTokensClaimed struct {
	ClaimConditionIndex *big.Int
	Claimer             common.Address
	Receiver            common.Address
	StartTokenId        *big.Int
	QuantityClaimed     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensClaimed is a free log retrieval operation binding the contract event 0xfa76a4010d9533e3e964f2930a65fb6042a12fa6ff5b08281837a10b0be7321e.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, address indexed claimer, address indexed receiver, uint256 startTokenId, uint256 quantityClaimed)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterTokensClaimed(opts *bind.FilterOpts, claimConditionIndex []*big.Int, claimer []common.Address, receiver []common.Address) (*MintAuraGeneratedTokensClaimedIterator, error) {

	var claimConditionIndexRule []interface{}
	for _, claimConditionIndexItem := range claimConditionIndex {
		claimConditionIndexRule = append(claimConditionIndexRule, claimConditionIndexItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "TokensClaimed", claimConditionIndexRule, claimerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedTokensClaimedIterator{contract: _MintAuraGenerated.contract, event: "TokensClaimed", logs: logs, sub: sub}, nil
}

// WatchTokensClaimed is a free log subscription operation binding the contract event 0xfa76a4010d9533e3e964f2930a65fb6042a12fa6ff5b08281837a10b0be7321e.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, address indexed claimer, address indexed receiver, uint256 startTokenId, uint256 quantityClaimed)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchTokensClaimed(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedTokensClaimed, claimConditionIndex []*big.Int, claimer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var claimConditionIndexRule []interface{}
	for _, claimConditionIndexItem := range claimConditionIndex {
		claimConditionIndexRule = append(claimConditionIndexRule, claimConditionIndexItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "TokensClaimed", claimConditionIndexRule, claimerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedTokensClaimed)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
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

// ParseTokensClaimed is a log parse operation binding the contract event 0xfa76a4010d9533e3e964f2930a65fb6042a12fa6ff5b08281837a10b0be7321e.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, address indexed claimer, address indexed receiver, uint256 startTokenId, uint256 quantityClaimed)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseTokensClaimed(log types.Log) (*MintAuraGeneratedTokensClaimed, error) {
	event := new(MintAuraGeneratedTokensClaimed)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintAuraGeneratedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MintAuraGenerated contract.
type MintAuraGeneratedTransferIterator struct {
	Event *MintAuraGeneratedTransfer // Event containing the contract specifics and raw log

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
func (it *MintAuraGeneratedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintAuraGeneratedTransfer)
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
		it.Event = new(MintAuraGeneratedTransfer)
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
func (it *MintAuraGeneratedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintAuraGeneratedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintAuraGeneratedTransfer represents a Transfer event raised by the MintAuraGenerated contract.
type MintAuraGeneratedTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MintAuraGeneratedTransferIterator, error) {

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

	logs, sub, err := _MintAuraGenerated.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MintAuraGeneratedTransferIterator{contract: _MintAuraGenerated.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MintAuraGenerated *MintAuraGeneratedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MintAuraGeneratedTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _MintAuraGenerated.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintAuraGeneratedTransfer)
				if err := _MintAuraGenerated.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MintAuraGenerated *MintAuraGeneratedFilterer) ParseTransfer(log types.Log) (*MintAuraGeneratedTransfer, error) {
	event := new(MintAuraGeneratedTransfer)
	if err := _MintAuraGenerated.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
