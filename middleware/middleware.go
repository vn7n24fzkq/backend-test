package middleware

import (
	"vn7n24fzkq/backend-test/common"

	"github.com/gin-gonic/gin"
)

type BaseMiddleware struct{}

func (p *BaseMiddleware) sendError(c *gin.Context, err *common.CommonError) {
	c.JSON(err.HttpCode, err)
}
