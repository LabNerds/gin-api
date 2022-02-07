package server

import (
	"gin-api/app/controllers"
	"gin-api/app/routers"
	"gin-api/helpers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(useLogger(helpers.DefaultLoggerFormatter), gin.Recovery())
	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	routers.V1(router)

	return router
}

func useLogger(formatter func(params gin.LogFormatterParams) string) gin.HandlerFunc {
	loggerConfig := helpers.NewLoggerConfig(formatter, []string{})
	return gin.LoggerWithConfig(loggerConfig)
}
