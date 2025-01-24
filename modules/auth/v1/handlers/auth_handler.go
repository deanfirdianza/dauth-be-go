package handler

import (
	"net/http"

	service "github.com/deanfirdianza/dauth-be-go/modules/auth/v1/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	// ...handle login logic...
	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func (h *AuthHandler) Register(c *gin.Context) {
	// ...handle register logic...
	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}
