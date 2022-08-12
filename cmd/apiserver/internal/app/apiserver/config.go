package apiserver

import "github.com/MaksimUlitin/api-server/cmd/apiserver/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_Addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

//New Config...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
