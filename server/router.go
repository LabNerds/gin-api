package server

import (
	"github.com/gin-gonic/gin"
	"gin-api/app/controllers"
	"gin-api/helpers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(helpers.LogFormatter))
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	return router
}
