package honey

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

var (
	c, _          = ethclient.Dial(os.Getenv("RPC_URL"))
	privateKey, _ = crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
)
