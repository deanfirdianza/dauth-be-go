package main

import (
	"log"

	"github.com/deanfirdianza/dauth-be-go/app/driver"
	"github.com/deanfirdianza/dauth-be-go/app/middlewares"
	authRoutes "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/routes"
	userRoutes "github.com/deanfirdianza/dauth-be-go/modules/users/v1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	conf := driver.Conf
	if driver.ErrConf != nil {
		log.Fatal(driver.ErrConf)
	}

	router.Use(middlewares.CORS())
	router.Use(middlewares.PanicRecovery)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authRoutes.AuthRoutes(router, driver.AuthHandler)
	userRoutes.UserRoutes(router, driver.UserHandler, conf, driver.DBSqlx)
	// AuthRoutes(router, conf, AuthHandler)

	router.Run(":" + conf.App.Port)
}
