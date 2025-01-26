package router

import (
	"cloud-martini-backend/handler"
	"cloud-martini-backend/queries"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/healthCheck", handler.HealthCheck)
	router.GET("/users", handler.GetUsers)
	router.POST("/users", func(ctx *gin.Context) {
		handler.AddUsers(ctx, queries.InsertUser)
	})
	router.DELETE("/user/:id", handler.DeleteUsers)
	router.PUT("/user/:id", handler.UpdateUsers)
	router.POST("/order")
	router.GET("/orders")

	return router
}
