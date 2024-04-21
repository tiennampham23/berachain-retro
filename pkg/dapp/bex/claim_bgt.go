package bex

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
)

var (
	ParsedRewardModuleABI, _ = RewardModuleMetaData.GetAbi()
)

func WithdrawAllDepositorRewards(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	poolAddress common.Address,
) (*types.Transaction, error) {
	//coins, err := CheckBGT(c, privateKey, poolAddress)
	//if err != nil {
	//	return nil, err
	//}
	//canClaimFunc := func(coin CosmosCoin) bool {
	//	return coin.Amount.Cmp(MinRewardAmountToClaim) > 0
	//}
	//
	//canClaim := false
	//for _, coin := range coins {
	//	if canClaimFunc(coin) {
	//		canClaim = true
	//	}
	//}
	//
	//if !canClaim {
	//	return nil, nil
	//}
	//
	claimRewardPayload, err := ParsedRewardModuleABI.Pack("withdrawAllDepositorRewards", poolAddress)
	if err != nil {
		return nil, err
	}

	log.Logger().Infow("Claim reward payload", "payload", hexutil.Encode(claimRewardPayload))

	tx, err := eth.SendTx(c, privateKey, claimRewardPayload, nil, RewardModuleContract)
	if err != nil {
		log.Logger().Errorw("Claim reward failed", "err", err)
		return nil, err
	}

	log.Logger().Infow("Claim reward successfully", "txHash", tx.Hash().Hex())
	return tx, nil
}

func CheckBGT(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	poolAddress common.Address,
) ([]CosmosCoin, error) {
	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)

	// function getCurrentRewards(address depositor, address receiver) external view returns (Cosmos.Coin[] memory);
	rewardModuleContract, err := NewRewardModuleCaller(RewardModuleContract, c)
	if err != nil {
		return nil, err
	}

	coins, err := rewardModuleContract.GetCurrentRewards(&bind.CallOpts{}, poolAddress, from)
	if err != nil {
		log.Logger().Errorw("Failed to get current rewards", "error", err)
		return nil, err
	}

	return coins, nil
}
