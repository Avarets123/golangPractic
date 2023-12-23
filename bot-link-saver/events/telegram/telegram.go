package event

import "bot-saver-v1/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	//storage
}

func New(client *telegram.Client) {
}
