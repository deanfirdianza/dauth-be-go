package handlers

import (
	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	// ...existing code...
}

func (h *UserHandler) Login(c *gin.Context) {
	// ...existing code...
}
