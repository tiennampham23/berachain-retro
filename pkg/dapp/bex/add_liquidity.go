package bex

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
)

func AddLiquidityForSingleToken(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	tokenIn common.Address,
	poolAddress common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)

	err := ApproveIfNeed(c, privateKey, tokenIn, amount)
	if err != nil {
		log.Logger().Errorw("Approve failed", "err", err)
		return nil, err
	}

	//  AddLiquidity(opts *bind.TransactOpts, pool common.Address, receiver common.Address, assetsIn []common.Address, amountsIn []*big.Int
	addLiquidityPayload, err := bexABI.Pack("addLiquidity", poolAddress, from, []common.Address{tokenIn}, []*big.Int{amount})
	if err != nil {
		log.Logger().Errorw("Pack addLiquidity failed", "err", err)
		return nil, err
	}

	log.Logger().Infow("Add Liquidity", "payload", hexutil.Encode(addLiquidityPayload))

	tx, err := eth.SendTx(c, privateKey, addLiquidityPayload, nil, SwapContractAddress)
	if err != nil {
		log.Logger().Errorw("Send AddLiquidity failed", "err", err)
		return nil, err
	}

	log.Logger().Infow("Send AddLiquidity tx successfully", "tx", tx.Hash().Hex())

	return tx, nil
}
