package models

import (
	"time"

	"github.com/google/uuid"
)

type Accounts struct {
	ID        uuid.UUID `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Salt      string    `db:"salt"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// ...existing code...
