package middleware

import (
	"net/http"
	"vn7n24fzkq/backend-test/common"

	"github.com/gin-gonic/gin"
)

type ErrorHandlerMiddleware struct {
	BaseMiddleware
}

func NewErrorHandlerMiddleware() *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{}
}

func (p *ErrorHandlerMiddleware) HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			switch err.Err.(type) {
			case nil:
				return
			case *common.CommonError:
				err := err.Err.(*common.CommonError)
				p.sendError(c, err)
			default:
				p.sendError(c, common.NewError(http.StatusInternalServerError, "UnknowError", err))
				return
			}
		}
	}
}
