package middleware

import (
	"errors"
	"net/http"
	"vn7n24fzkq/backend-test/common"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	BaseMiddleware
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (p *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetHeader("Authorization")
		var err error = nil
		if jwt == "" {
			c.Abort()
			p.sendError(c,
				common.NewError(
					http.StatusUnauthorized, "You might forgot put JWT into Authorization header",
					errors.New("Authorization header should not be empty"),
				),
			)
			return
		}

		jwtUser, err := common.DecodeJWT(jwt)
		if err != nil {
			c.Abort()
			p.sendError(c, common.NewError(http.StatusUnauthorized, "JWT validate failed", err))
			return
		}

		p.setCurrentUser(c, &jwtUser)

		c.Next()
	}
}

func (p *AuthMiddleware) setCurrentUser(c *gin.Context, jwtUser *common.JWTUser) {
	c.Set("JWTUser", jwtUser)
}

// Only authorized request can use this function, owtherwise it panic
func (p *AuthMiddleware) GetCurrentUser(c *gin.Context) *common.JWTUser {
	return c.MustGet("JWTUser").(*common.JWTUser)
}
