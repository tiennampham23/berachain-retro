package config

import (
	"encoding/json"
	"os"
)

func ReadConfig(relativePath string) (*Config, error) {
	file, err := os.Open(relativePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config

	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
