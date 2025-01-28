package models

type AuthRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
