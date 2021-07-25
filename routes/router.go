package routes

import "github.com/gin-gonic/gin"
import "vn7n24fzkq/backend-test/routes/api"

func InitRouter() *gin.Engine {
	gin := gin.Default()
	// register route
	registerRoutes(gin)
	gin.Run(":8080")
	return gin
}

func registerRoutes(server *gin.Engine) {
	server.GET("/users/:id", api.GetUserById)
}
