package server

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers"
	"gin-api/utils"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(utils.LogFormatter))
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	return router
}
