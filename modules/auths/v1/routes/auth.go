package routes

import (
	handler "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authHandler *handler.AuthHandler) {
	// AuthRoutes for authentication

	authRoutes := router.Group("/v1/auth")
	authRoutes.POST("/login", authHandler.Login)
	authRoutes.POST("/register", authHandler.Register)
	authRoutes.GET("/validate-token", authHandler.ValidateJWT)
	authRoutes.GET("/refresh-token", authHandler.RefreshToken)
	authRoutes.POST("/logout", authHandler.Logout)

}
