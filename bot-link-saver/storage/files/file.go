package files

import (
	"bot-saver-v1/lib/customErrors"
	"bot-saver-v1/storage"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	basePath string
}

const basePerm = 0774

var ErrNoSavedPages = errors.New("No saved page")

func New(basePath string) *Storage {
	return &Storage{basePath: basePath}
}

func (s *Storage) Save(page *storage.Page) (err error) {
	defer func() {
		err = customErrors.WrapIfErr("Can't save save", err)
	}()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, basePerm); err != nil {
		return err
	}

	fileName, err := fileName(page)

	if err != nil {
		return err
	}

	filePath = filepath.Join(filePath, fileName)

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil

}

func (s *Storage) PickRandom(userName string) (page *storage.Page, err error) {

	defer func() {
		err = customErrors.WrapIfErr("Can't save save", err)
	}()

	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var filesCount = len(files)

	if filesCount == 0 {
		return nil, ErrNoSavedPages
	}

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(filesCount)

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))

}

func (s *Storage) Remove(p *storage.Page) error {

	fileName, err := fileName(p)

	if err != nil {
		return customErrors.Wrap("Can't remove file: "+string(fileName), err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err != nil {
		return customErrors.Wrap("Can't remove file: "+string(fileName), err)
	}

	return nil

}

func (s *Storage) IsExists(p *storage.Page) (bool, error) {

	fileName, err := fileName(p)

	if err != nil {
		return false, customErrors.Wrap("Can't get file: "+string(fileName), err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil

	case err != nil:
		msg := fmt.Sprintf("Can't check if file %s exists", path)
		return false, customErrors.Wrap(msg, err)
	}

	return true, nil

}

func (s *Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, customErrors.Wrap("Can't decode page", err)
	}

	defer func() { _ = f.Close() }()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, customErrors.Wrap("Can't decode page", err)
	}

	return &p, nil

}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
