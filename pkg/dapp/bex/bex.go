package bex

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"math/big"
	"time"
)

var (
	WBear                        = common.HexToAddress("0x5806E416dA447b267cEA759358cF22Cc41FAE80F")
	Honey                        = common.HexToAddress("0x7EeCA4205fF31f947EdBd49195a7A88E6A91161B")
	Bear                         = common.HexToAddress("0x0000000000000000000000000000000000000000")
	Erc20ModuleContract          = common.HexToAddress("0x0000000000000000000000000000000000696969")
	bexABI, _                    = BexGeneratedMetaData.GetAbi()
	batchSwapKind          uint8 = 0
	ApproveMax, _                = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
	SwapContractAddress          = common.HexToAddress("0x0d5862FDbdd12490f9b4De54c236cff63B038074")
	RewardModuleContract         = common.HexToAddress("0x55684e2cA2bace0aDc512C1AFF880b15b8eA7214")
	MinRewardAmountToClaim       = new(big.Int).SetInt64(0)
)

var (
	client *resty.Client
)

func init() {
	client = resty.New().SetTimeout(10 * time.Second)
}
