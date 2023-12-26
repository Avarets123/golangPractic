package postgres

import (
	"v1/pkg/storage"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

func NewPostgresDb(connectionStr string, log *logrus.Logger) *storage.SqlStorage {

	storeConfig := storage.SqlStorageConfig{
		Type:          "postgres",
		ConnectionStr: connectionStr,
	}

	return storage.NewSqlStorage(&storeConfig, log)

}
