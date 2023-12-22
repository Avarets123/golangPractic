package main

import (
	"bot-saver-v1/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	_ = telegram.New(tgBotHost, mustToken())

	//

}

func mustToken() string {

	token := flag.String("bot-token", "", "telegram bot token")

	if *token == "" {
		log.Fatal("Telegram token not passed")

	}

	return *token

}
