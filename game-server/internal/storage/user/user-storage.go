package user_storage

import (
	"database/sql"
	"v1/internal/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserStorage struct {
	db     *sql.DB
	logger *logrus.Logger
}

type UserStorageInterface interface {
	SaveUser(nickname, password string) (string, error)
	FindOne(id string) *domain.User
	FindMany(limit, offset int) *[]domain.User
	Delete(id string)
}

func New(db *sql.DB, logger *logrus.Logger) *UserStorage {

	return &UserStorage{db: db, logger: logger}
}

func (us *UserStorage) SaveUser(nickname, password string) (id string, retErr error) {

	var err error

	defer func() { retErr = err }()

	stmt, err := us.db.Prepare(`INSERT INTO users (id, nickname, password) VALUES ($1, $2, $3) RETURNING id`)

	id = uuid.New().String()

	var newId string

	err = stmt.QueryRow(id, nickname, password).Scan(&newId)

	defer stmt.Close()

	us.logger.Info("UUID is " + id)
	us.logger.Info("new UUID is " + newId)

	return

}
