package main

import (
	"flag"
	"v1/internal/apiserver"
	"v1/internal/config"
	configureLogrus "v1/pkg/logger/logrus"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func main() {

	flag.StringVar(&configPath, "config-path", "configs/dev.config.yml", "path to config file")
	flag.Parse()

	config := config.NewConfig()

	cleanenv.ReadConfig(configPath, &config)

	logger := initLogrus()

	server := apiserver.New(config, logger)

	err := server.Start()

	if err != nil {
		panic(err)
	}

}

func initLogrus() *logrus.Logger {

	logrusConfig := configureLogrus.LogrusConfig{
		OutputType: configureLogrus.TEXT,
		LogLevel:   "INFO",
	}

	logrus := configureLogrus.SetUpLogrus(&logrusConfig)

	return logrus

}
