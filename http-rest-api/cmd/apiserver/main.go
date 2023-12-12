package main

import (
	"flag"
	"log"
	"rest-api/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to file config")
}

func main() {

	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	errorHandler(err)

	s := apiserver.New(config)
	err = s.Start()
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
