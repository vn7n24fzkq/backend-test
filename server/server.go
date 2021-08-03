package server

import (
	"vn7n24fzkq/backend-test/dao"
	"vn7n24fzkq/backend-test/middleware"
	"vn7n24fzkq/backend-test/routes/api"
	"vn7n24fzkq/backend-test/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	gin *gin.Engine
}

func InstanceServer(engine *gin.Engine, db *gorm.DB) *Server {
	// Initialize DAO
	userDAO := dao.NewUserDAO(db)
	taskDAO := dao.NewTaskDAO(db)

	// Initialize Service
	userService := service.NewUserService(userDAO)
	taskService := service.NewTaskService(taskDAO, userService)

	// Initialize Middleware
	authMiddleware := &middleware.AuthMiddleware{}
	errorHandlerMiddleware := &middleware.ErrorHandlerMiddleware{}
	apiRoutes := api.NewAPIRouter(
		errorHandlerMiddleware,
		authMiddleware,
		userService,
		taskService,
	)

	server := &Server{
		gin: engine,
	}

	// register route
	server.registerRoutes(apiRoutes)
	return server
}

func (p *Server) Run(addr string) {
	p.gin.Run(addr)
}

func (p *Server) registerRoutes(apiRouter *api.APIRouter) {
	p.gin.Static("/static", "./static")
	api := p.gin.Group("/api", apiRouter.ErrorHandler.HandleError())
	{
		api.POST("/registration", apiRouter.Registration)
		api.POST("/auth", apiRouter.Auth)

		authAPI := api.Group("", apiRouter.AuthMiddleware.Auth())
		{
			authAPI.GET("/users/:id", apiRouter.GetUserById)

			authAPI.GET("/tasks", apiRouter.GetAllTasks)
			authAPI.POST("/tasks", apiRouter.CreateTask)
			authAPI.PUT("/tasks/:id", apiRouter.UpdateTask)
			authAPI.DELETE("/tasks/:id", apiRouter.DeleteTask)
		}
	}

}
