package server

import (
	"gin-api/app/controllers"
	"gin-api/app/routers"
	"gin-api/helpers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(helpers.LogFormatter))
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	routers.V1(router)

	return router
}
