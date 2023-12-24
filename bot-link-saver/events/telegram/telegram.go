package telegram

import (
	"bot-saver-v1/clients/telegram"
	"bot-saver-v1/events"
	"bot-saver-v1/lib/customErrors"
	"bot-saver-v1/storage"
	"errors"
)

var ErrUnknownEventType = errors.New("unknown error type")

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	Username string
}

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {

	updates, err := p.tg.Updates(p.offset, limit)

	if err != nil {
		return nil, customErrors.Wrap("Can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil

}

func (p *Processor) Proccess(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)

	default:
		return customErrors.Wrap("cannot process event", ErrUnknownEventType)
	}
}

func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)

	if err != nil {
		return customErrors.Wrap("Can't process message", err)
	}

	if err := p.doCmd(event.Text, meta.Username, meta.ChatID); err != nil {
		return customErrors.Wrap("Can't process message", err)
	}

	return nil

}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)

	if !ok {
		return Meta{}, customErrors.Wrap("Can't get meta", errors.New("Can't get meta"))
	}

	return res, nil
}

func event(u telegram.Update) events.Event {

	uType := fetchType(u)

	res := events.Event{
		Type: uType,
		Text: fetchText(u),
	}

	if uType == events.Message {
		res.Meta = Meta{
			ChatID:   u.Message.Chat.ID,
			Username: u.Message.From.Username,
		}
	}

	return res
}

func fetchText(u telegram.Update) string {
	if u.Message == nil {
		return ""
	}

	return u.Message.Text
}

func fetchType(u telegram.Update) events.Type {
	if u.Message == nil {
		return events.Unknown
	}

	return events.Message
}
