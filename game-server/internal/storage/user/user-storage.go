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
	SaveUser(nickname, password string) (id string, retErr error)
	FindOne(id string) (user *domain.User, retErr error)
	FindMany(limit, offset int) (users *[]domain.User, retErr error)
	Delete(id string) (retErr error)
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

	return
}

func (us *UserStorage) FindOne(id string) (user *domain.User, retErr error) {

	var err error

	defer func() { retErr = err }()

	stmt, err := us.db.Prepare(`SELECT * FROM users WHERE id = $1`)

	err = stmt.QueryRow(id).Scan(user)

	return

}

func (us *UserStorage) FindMany(limit, offset int) (users []*domain.User, retErr error) {

	var err error

	defer func() { retErr = err }()

	rows, err := us.db.Query(`SELECT * FROM users LIMIT $1 OFFSET $2`, limit, offset)

	for rows.Next() {

		user := domain.User{}

		err = rows.Scan(&user.ID, &user.Nickname, &user.Password, &user.CreatedAt)

		users = append(users, &user)

	}

	return

}

func (us *UserStorage) Delete(id string) (retErr error) {

	var err error

	defer func() { retErr = err }()

	stmt, err := us.db.Prepare(`DELETE FROM users WHERE id = $1`)

	_, err = stmt.Exec(id)

	return

}
