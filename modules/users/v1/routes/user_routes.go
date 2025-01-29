package routes

import (
	"github.com/deanfirdianza/dauth-be-go/app/env"
	"github.com/deanfirdianza/dauth-be-go/app/middlewares"
	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func UserRoutes(router *gin.Engine, userHandler *handlers.UserHandler, conf env.Conf, DBSqlx *sqlx.DB) {
	auth := middlewares.Auth(conf, DBSqlx)
	userGroup := router.Group("v1/user")
	{
		userGroup.GET("/profile", auth, userHandler.Profile)
		userGroup.POST("/", auth, userHandler.Register)
		userGroup.DELETE("/", auth, userHandler.Register)
	}
}
