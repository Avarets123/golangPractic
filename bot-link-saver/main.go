package main

import (
	tgClient "bot-saver-v1/clients/telegram"
	event_consumer "bot-saver-v1/consumer/event-consumer"
	"bot-saver-v1/events/telegram"
	"bot-saver-v1/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 15
)

func main() {

	newTgClient := tgClient.New(tgBotHost, mustToken())

	eventsProcessor := telegram.New(newTgClient, files.New(storagePath))

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Service is stopped", err)
	}

}

func mustToken() string {

	token := flag.String("bot-token", "", "telegram bot token")
	flag.Parse()

	if *token == "" {
		log.Fatal("Telegram token not passed")

	}

	return *token

}
