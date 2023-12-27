package domain

import "time"

type User struct {
	ID,
	Nickname,
	Password string
	CreatedAt time.Time
}
