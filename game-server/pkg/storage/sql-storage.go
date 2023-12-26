package storage

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type SqlStorage struct {
	Db     *sql.DB
	Logger *logrus.Logger
}

type SqlStorageConfig struct {
	Type,
	ConnectionStr string
}

func NewSqlStorage(conf *SqlStorageConfig, log *logrus.Logger) *SqlStorage {

	db, err := sql.Open(conf.Type, conf.ConnectionStr)

	if err != nil {
		log.Fatal("Error in initing db!")
	}

	log.Info("Connection to db has been succesfully completed")

	return &SqlStorage{
		Db:     db,
		Logger: log,
	}

}

func (s *SqlStorage) Connect() {}
