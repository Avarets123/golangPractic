package domain

import "time"

type User struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
