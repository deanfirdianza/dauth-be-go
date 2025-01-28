package routes

import (
	handler "github.com/deanfirdianza/dauth-be-go/modules/auth/v1/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authHandler *handler.AuthHandler) {
	// AuthRoutes for authentication

	authRoutes := router.Group("/v1/auth")
	authRoutes.POST("/login", authHandler.Login)
	authRoutes.POST("/register", authHandler.Register)
	authRoutes.GET("/refresh-token", authHandler.Register)
	authRoutes.POST("/logout", authHandler.Login)

}
