package models

import (
	"time"
)

// ...existing code...

type User struct {
	ID        uint      `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Salt      string    `db:"salt"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// ...existing code...
