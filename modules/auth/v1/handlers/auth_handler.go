package handler

import (
	"fmt"
	"net/http"

	authModel "github.com/deanfirdianza/dauth-be-go/modules/auth/v1/models"
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
	var register authModel.AuthRegister
	err := c.ShouldBindBodyWithJSON(&register)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(register.Email)
	err = h.authService.Register(register.Username, register.Password, register.Email)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}
