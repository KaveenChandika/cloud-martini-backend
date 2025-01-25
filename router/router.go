package router

import (
	"cloud-martini-backend/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/healthCheck", handler.HealthCheck)
	router.GET("/users", handler.GetUsers)
	router.POST("/users", handler.AddUsers)
	router.POST("/order")
	router.GET("/orders")

	return router
}
