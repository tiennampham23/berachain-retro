package bex

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

/***
{
    "chainId": 80085,
    "data": "0xd0e30db0",
    "from": "0x3122C65dA0E288Fb745F07D8C81B10427b28E7aD",
    "gas": "0xd0fd",
    "gasPrice": "0xb38bcae2",
    "nonce": "0x1a",
    "to": "0x5806E416dA447b267cEA759358cF22Cc41FAE80F",
    "value": "0x2386f26fc10000"
}
*/
/**
to = 0x5806E416dA447b267cEA759358cF22Cc41FAE80F
*/
/**
https://artio.beratrail.io/tx/0x1a91232ad62533c084eb7a68c230ae7518799f6b87dcb5608666297ec8f1c3e3
*/

func Wrap() ([]byte, error) {
	return hexutil.Decode("0xd0e30db0") // deposit()
}
