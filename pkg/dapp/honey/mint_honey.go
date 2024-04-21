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
	HoneyContractAddress = common.HexToAddress("0x09ec711b81cD27A6466EC40960F2f8D85BB129D9")
	STGUSDC              = common.HexToAddress("0x6581e59A1C8dA66eD0D313a0d4029DcE2F746Cc5")
	honeyABI, _          = HoneyGeneratedMetaData.GetAbi()
)

/** approve
{
	"chainId": 80085,
	"data": "0x095ea7b300000000000000000000000009ec711b81cd27a6466ec40960f2f8d85bb129d97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	"from": "0x3122C65dA0E288Fb745F07D8C81B10427b28E7aD",
	"gas": "0x115c7",
	"gasPrice": "0x4031755",
	"nonce": "0x23",
	"to": "0x6581e59A1C8dA66eD0D313a0d4029DcE2F746Cc5",
	"value": "0x0"
}

*/

/** mint
{
  "chainId": 80085,
  "data": "0xc6c3bbe60000000000000000000000003122c65da0e288fb745f07d8c81b10427b28e7ad0000000000000000000000006581e59a1c8da66ed0d313a0d4029dce2f746cc5000000000000000000000000000000000000000000000000a12f09c0afad9b5a",
  "from": "0x3122C65dA0E288Fb745F07D8C81B10427b28E7aD",
  "gas": "0x388cb",
  "gasPrice": "0x391e305",
  "nonce": "0x24",
  "to": "0x09ec711b81cD27A6466EC40960F2f8D85BB129D9",
  "value": "0x0"
}
*/

func Mint(
	c *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	tokenIn common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	err := dappCommon.ApproveIfNeed(c, privateKey, tokenIn, amount, HoneyContractAddress)
	if err != nil {
		return nil, err
	}

	var (
		from, _ = eth.GetAddressFromPrivateKey(privateKey)
	)

	payload, err := honeyABI.Pack("mint", from, tokenIn, amount)
	if err != nil {
		return nil, err
	}

	log.Logger().Infow("Mint payload", "payload", hexutil.Encode(payload))

	tx, err := eth.SendTx(c, privateKey, payload, nil, HoneyContractAddress)
	if err != nil {
		log.Logger().Errorw("failed to send mint", "error", err)
		return nil, err
	}

	return tx, nil
}
