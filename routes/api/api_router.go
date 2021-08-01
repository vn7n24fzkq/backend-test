package api

import (
	"vn7n24fzkq/backend-test/common"
	"vn7n24fzkq/backend-test/middleware"
	"vn7n24fzkq/backend-test/routes"
	"vn7n24fzkq/backend-test/service"

	"github.com/gin-gonic/gin"
)

type APIRouter struct {
	AuthMiddleware *middleware.AuthMiddleware
	UserService    *service.UserService
	TaskService    *service.TaskService
	*routes.BaseRouter
}

func NewAPIRouter(
	errorHandlerMiddleware *middleware.ErrorHandlerMiddleware,
	authMiddleware *middleware.AuthMiddleware,
	userService *service.UserService,
	taskService *service.TaskService,
) *APIRouter {
	return &APIRouter{
		AuthMiddleware: authMiddleware,
		UserService:    userService,
		TaskService:    taskService,
		BaseRouter: &routes.BaseRouter{
			ErrorHandler: errorHandlerMiddleware,
		},
	}
}

func (p *APIRouter) GetCurrentUser(c *gin.Context) *common.JWTUser {
	return p.AuthMiddleware.GetCurrentUser(c)
}
