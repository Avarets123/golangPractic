package telegram

import (
	"bot-saver-v1/clients/telegram"
	"bot-saver-v1/lib/customErrors"
	"bot-saver-v1/storage"
	"bot-saver-v1/storage/files"
	"errors"
	"log"
	"net/url"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text, username string, chatID int) error {

	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.SendRandom(chatID, username)
	case HelpCmd:
		return p.SendHelp(chatID)
	case StartCmd:
		return p.SendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)

	}

}

func (p *Processor) savePage(chatID int, pageUrl, username string) (err error) {

	defer func() { err = customErrors.WrapIfErr("Can't do command: save page", err) }()

	page := &storage.Page{
		URL:      pageUrl,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(page)

	if err != nil {
		return err
	}
	var messageSender = wrapperSendMessage(chatID, p.tg)

	if isExists {
		return messageSender(msgAlreadyExists)
		// return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil {
		return err
	}

	if err := messageSender(msgSaved); err != nil {
		return err
	}

	return nil

}

func (p *Processor) SendRandom(chatID int, username string) (err error) {
	defer func() { err = customErrors.WrapIfErr("Can't do command: send random", err) }()

	page, err := p.storage.PickRandom(username)

	if err != nil && !errors.Is(err, files.ErrNoSavedPages) {
		return err
	}

	var messageSender = wrapperSendMessage(chatID, p.tg)

	if errors.Is(err, files.ErrNoSavedPages) {
		return messageSender(msgNoSavedPages)
	}

	if err := messageSender(page.URL); err != nil {
		return err
	}

	return p.storage.Remove(page)

}

func (p *Processor) SendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) SendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func wrapperSendMessage(chatID int, tg *telegram.Client) func(string) error {
	return func(msg string) error {
		return tg.SendMessage(chatID, msg)
	}

}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
