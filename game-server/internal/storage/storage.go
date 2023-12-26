package storage2

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Storage struct {
	Db *sql.DB
}

type LogInfo interface {
	Info(args ...interface{})
}

func New(log *logrus.Logger) (storage *Storage, err error) {

	const connectionStr = "postgresql://postgres:password@localhost:5434/postgres?schema=public"

	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		return
	}

	log.Info("Connection to db has been succesfully completed")

	storage = &Storage{
		Db: db,
	}

	return

}
