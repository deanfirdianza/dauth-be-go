package handlers

import (
	"fmt"
	"net/http"

	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Profile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	fmt.Println(userID)
	user, err := h.userService.GetUserDetail(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("user :", user)
	c.JSON(http.StatusOK, gin.H{"data": user})
	// ...existing code...
}

func (h *UserHandler) Register(c *gin.Context) {
	// ...existing code...
}

func (h *UserHandler) Login(c *gin.Context) {
	// ...existing code...
}
