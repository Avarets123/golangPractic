package event_consumer

import (
	"bot-saver-v1/events"
	"log"
	"time"
)

type Consumer struct {
	feetcher  events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) *Consumer {
	return &Consumer{
		feetcher:  fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		newEvents, err := c.feetcher.Fetch(c.batchSize)

		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())
			continue
		}

		if len(newEvents) == 0 {
			time.Sleep(time.Second)

			continue
		}

		if err := c.handleEvents(newEvents); err != nil {
			log.Print(err)

			continue
		}

	}
}

func (c Consumer) handleEvents(events []events.Event) error {

	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Proccess(event); err != nil {
			log.Printf("can't handle event: %s", event.Text)

			continue

		}

	}

	return nil

}
