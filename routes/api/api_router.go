package api

import "vn7n24fzkq/backend-test/service"

type APIRoutes struct {
	UserService *service.UserService
	TaskService *service.TaskService
}
