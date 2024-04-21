package honey

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	dappCommon "github.com/tiennampham23/berachain-airdrop/pkg/dapp/common"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
)

var (
	HoneyToken = common.HexToAddress("0x7EeCA4205fF31f947EdBd49195a7A88E6A91161B")
)

func RedeemHoney(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	tokenOut common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	err := dappCommon.ApproveIfNeed(c, privateKey, HoneyToken, amount, HoneyContractAddress)
	if err != nil {
		return nil, err
	}

	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)

	//  Redeem(opts *bind.TransactOpts, to common.Address, amount *big.Int, tokenOut common.Address)
	// 	return _HoneyGenerated.contract.Transact(opts, "redeem", to, amount, tokenOut)
	payload, err := honeyABI.Pack("redeem", from, amount, tokenOut)
	if err != nil {
		return nil, err
	}

	log.Logger().Infow("Redeem payload", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, nil, HoneyContractAddress)
	if err != nil {
		log.Logger().Errorw("failed to send mint", "error", err)
		return nil, err
	}

	return tx, nil

}
