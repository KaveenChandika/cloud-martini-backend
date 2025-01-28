package router

import (
	"cloud-martini-backend/handler"
	"cloud-martini-backend/queries"
	"cloud-martini-backend/router/middleware.go"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.GET("/health", handler.HealthCheck)
	router.GET("/users", func(ctx *gin.Context) {
		handler.GetUsers(ctx, queries.GetUsers)
	})
	router.POST("/users", func(ctx *gin.Context) {
		handler.AddUsers(ctx, queries.InsertUser)
	})
	router.DELETE("/user/:id", func(ctx *gin.Context) {
		handler.DeleteUsers(ctx, queries.DeleteUser)
	})
	router.PUT("/user/:id", func(ctx *gin.Context) {
		handler.UpdateUsers(ctx, queries.UpdateUsers)
	})
	router.POST("/order")
	router.GET("/orders")

	return router
}
