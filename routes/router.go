package routes

import (
	"github.com/gin-gonic/gin"
	"vn7n24fzkq/backend-test/routes/api"
)

func InitRouter() *gin.Engine {
	gin := gin.Default()
	// register route
	registerRoutes(gin)
	gin.Run(":8080")
	return gin
}

func registerRoutes(server *gin.Engine) {

	server.POST("/api/auth/registration", api.Registration)
	server.POST("/api/auth", api.Auth)

	server.GET("/api/users/:id", api.GetUserById)

	server.GET("/api/tasks", api.GetAllTasks)
	server.GET("/api/tasks/:id", api.GetTaskById)
	server.POST("/api/tasks", api.CreateTask)
	server.PUT("/api/tasks/:id", api.UpdateTask)
	server.DELETE("/api/tasks/:id", api.DeleteTask)
}
