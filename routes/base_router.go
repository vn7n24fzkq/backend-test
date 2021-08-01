package routes

import "vn7n24fzkq/backend-test/middleware"

type BaseRouter struct {
	ErrorHandler *middleware.ErrorHandlerMiddleware
}
