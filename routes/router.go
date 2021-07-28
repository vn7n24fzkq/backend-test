package routes

import (
	"vn7n24fzkq/backend-test/routes/api"
	"vn7n24fzkq/backend-test/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
}

func InstanceServer(engine *gin.Engine, userService *service.UserService, taskService *service.TaskService) *Server {
	apiRoutes := &api.APIRoutes{UserService: userService, TaskService: taskService}
	server := &Server{gin: engine}
	// register route
	server.registerRoutes(apiRoutes)
	return server
}

func (p *Server) Run(addr string) {
	p.gin.Run(addr)
}

func (p *Server) registerRoutes(apiRouter *api.APIRoutes) {
	p.gin.POST("/api/auth/registration", apiRouter.Registration)
	p.gin.POST("/api/auth", apiRouter.Auth)

	p.gin.GET("/api/users/:id", apiRouter.GetUserById)

	p.gin.GET("/api/tasks", apiRouter.GetAllTasks)
	p.gin.GET("/api/tasks/:id", apiRouter.GetTaskById)
	p.gin.POST("/api/tasks", apiRouter.CreateTask)
	p.gin.PUT("/api/tasks/:id", apiRouter.UpdateTask)
	p.gin.DELETE("/api/tasks/:id", apiRouter.DeleteTask)
}
