package main

import (
	"github.com/gin-gonic/gin"
	"gin-api/src/utils"
)

func main() {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(utils.LogFormatter))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
