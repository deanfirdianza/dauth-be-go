package main

import (
	"github.com/deanfirdianza/dauth-be-go/app/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.CORS())
	router.Use(middlewares.PanicRecovery)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
