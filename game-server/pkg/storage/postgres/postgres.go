package postgres

import (
	"database/sql"
	"v1/pkg/storage"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

func NewPostgresDb(connectionStr string, log *logrus.Logger) *storage.SqlStorage {

	storeConfig := storage.SqlStorageConfig{
		Type:          "postgres",
		ConnectionStr: connectionStr,
	}

	db := storage.NewSqlStorage(&storeConfig, log)

	createUsersTable(db.Db, log)

	return db

}

func createUsersTable(db *sql.DB, logger *logrus.Logger) {
	_, err := db.Query(`CREATE TABLE IF NOT EXISTS "users" (
		"id" UUID primary key,
		"nickname" varchar(40) not null,
		"password" varchar(100) not null,
		"createdAt" TIMESTAMP(3) not null DEFAULT CURRENT_TIMESTAMP
	);
	CREATE UNIQUE INDEX IF NOT EXISTS "users_nickname_unique" on "users"("nickname");
	
	`)

	if err != nil {
		logger.Fatal(err)
	}

}
