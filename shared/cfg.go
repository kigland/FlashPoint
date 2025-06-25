package shared

import "encoding/json"

type Config struct {
	Addr  string `json:"addr"`
	Debug bool   `json:"debug"`

	APIKeys []string `json:"api_keys"`
}

var cfg *Config

func GetConfig() *Config {
	return cfg
}

func LoadConfig(bs []byte) error {
	return json.Unmarshal(bs, &cfg)
}
