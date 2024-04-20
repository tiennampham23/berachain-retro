package bot

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tiennampham23/berachain-airdrop/pkg/config"
	"github.com/tiennampham23/berachain-airdrop/pkg/dapp/mintaura"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/tiennampham23/berachain-airdrop/pkg/utilties/eth"
	"math/big"
)

type Bot struct {
	ethClient *ethclient.Client
	cfg       *config.Config
}

func NewBot(relativePath string) (*Bot, error) {
	cfg, err := config.ReadConfig(relativePath)
	if err != nil {
		log.Logger().Fatalw("Failed to read config", "err", err)
	}

	ethClient, err := ethclient.Dial(cfg.Rpc)
	if err != nil {
		return nil, err
	}

	return &Bot{
		ethClient: ethClient,
		cfg:       cfg,
	}, nil
}

func (b *Bot) Start() error {
	privateKeys, err := b.getPrivateKeys()
	if err != nil {
		return err
	}

	for _, privateKey := range privateKeys {
		log.Logger().Infow("Private key", "privateKey", privateKey)
	}
	return nil
}

func (b *Bot) getPrivateKeys() ([]*ecdsa.PrivateKey, error) {
	privateKeys := make([]*ecdsa.PrivateKey, 0)
	for _, key := range b.cfg.PrivateKeys {
		privateKey, err := crypto.HexToECDSA(key)
		if err != nil {
			return nil, err
		}
		privateKeys = append(privateKeys, privateKey)
	}
	return privateKeys, nil
}

func (b *Bot) claimNFTMintAura(fromAddress common.Address, privateKey *ecdsa.PrivateKey) error {
	quantity := big.NewInt(1)
	currency := common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	pricePerToken := big.NewInt(50000000000000000)

	quantityLimitPerWallet, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	byte32 := [32]byte{}
	allowlistProof := mintaura.IDropAllowlistProof{
		Proof:                  [][32]byte{byte32},
		QuantityLimitPerWallet: quantityLimitPerWallet,
		PricePerToken:          pricePerToken,
		Currency:               currency,
	}

	var data []byte

	payload, err := mintaura.Claim(
		fromAddress,
		quantity,
		currency,
		pricePerToken,
		allowlistProof,
		data,
	)

	if err != nil {
		log.Logger().Errorw("Failed to claim NFT Mint Aura", "err", err, "address", fromAddress)
		return err
	}

	to := common.HexToAddress("0x6Dc544b917904959e5c26b9202907ca16BDcdb4A")

	tx, err := eth.SendTx(b.ethClient, privateKey, payload, pricePerToken, to)
	if err != nil {
		log.Logger().Errorw("Failed to claim NFT Mint Aura", "err", err, "address", fromAddress)
		return err
	}

	log.Logger().Infow("Mint NFT Aura successfully", "tx", tx.Hash().Hex())
	return nil
}
