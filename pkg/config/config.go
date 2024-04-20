package config

type SwapBex struct {
	Enable   bool   `json:"enable" yaml:"enable"`
	TokenIn  string `json:"token_in" yaml:"token_in"`
	TokenOut string `json:"token_out" yaml:"token_out"`
}

type Config struct {
	Rpc            string   `json:"rpc" yaml:"rpc"`
	PrivateKeys    []string `json:"private_keys" yaml:"private_keys"`
	EnableMintAura bool     `json:"enable_mint_aura" yaml:"enable_mint_aura"`
	EnableWrapBera bool     `json:"enable_wrap_bera" yaml:"enable_wrap_bera"`
	SwapBex        SwapBex  `json:"swap_bex" yaml:"swap_bex"`
}
