package models

import "time"

type AuthRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthToken struct {
	ID           int       `json:"id" db:"id"`
	AccountID    string    `json:"account_id" db:"account_id"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
	Revoked      bool      `json:"revoked" db:"revoked"`
}
