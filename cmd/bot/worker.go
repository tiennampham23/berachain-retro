package bot

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tiennampham23/berachain-airdrop/pkg/config"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
)

const (
	NativeToken = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
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
		// Interact with dapps
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
