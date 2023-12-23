package storage

import (
	"bot-saver-v1/lib/customErrors"
	"crypto/sha512"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL,
	UserName string
}

func (p *Page) Hash() (string, error) {

	h := sha512.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", customErrors.Wrap("Can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", customErrors.Wrap("Can't calculate hash", err)
	}

	return string(h.Sum(nil)), nil

}
