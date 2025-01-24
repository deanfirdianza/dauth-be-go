package routes

import (
	"github.com/deanfirdianza/dauth-be-go/app/middlewares"
	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	auth := middlewares.Auth()
	userGroup := router.Group("/user")
	{
		userGroup.GET("/profile", auth, userHandler.Register)
		userGroup.POST("/", auth, userHandler.Register)
		userGroup.DELETE("/", auth, userHandler.Register)
	}
}
