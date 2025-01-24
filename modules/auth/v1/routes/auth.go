package routes

import (
	"github.com/deanfirdianza/dauth-be-go/app/env"
	handler "github.com/deanfirdianza/dauth-be-go/modules/v1/auth/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, conf *env.Conf, authHandler handler.AuthHandler) {
	// AuthRoutes for authentication

	// authRoutes := router.Group("/auth")
	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)

}
