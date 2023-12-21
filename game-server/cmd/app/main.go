package main

import (
	"flag"
	"v1/internal/app/apiserver"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	configPath string
)

func main() {

	flag.StringVar(&configPath, "config-path", "configs/dev.config.yml", "path to config file")
	flag.Parse()

	config := apiserver.NewConfig()

	cleanenv.ReadConfig(configPath, &config)

	server := apiserver.New(config)

	err := server.Start()

	if err != nil {
		panic(err)
	}

}
