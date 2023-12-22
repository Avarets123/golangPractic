package users

import (
	"database/sql"

	"github.com/google/uuid"
)

type Users struct {
	db *sql.DB
}

type userRow struct {
	id string
}

func New(db *sql.DB) *Users {

	return &Users{
		db: db,
	}

}

func (u *Users) Save(emal, password string) string {

	id := uuid.New().String()

	stmt, err := u.db.Prepare(`INSERT INTO USERS ("id", "email", "password") VALUES(?, ?, ?) RETURNING id`)

	if err != nil {
		panic(err)
	}

	// res, err := stmt.Exec(id, emal, password)

	res := stmt.QueryRow(id, emal, password)

	user := userRow{}

	res.Scan(user)

	return user.id

}
