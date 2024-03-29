package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/MaksimUlitin/api-server/cmd/apiserver/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "paht to config file ")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	s := apiserver.New(config)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
