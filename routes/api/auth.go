package api

import (
	"net/http"
	"vn7n24fzkq/backend-test/common"

	"github.com/gin-gonic/gin"
)

type AuthResponse struct {
	JWT string `json:"jwt"`
}

type AuthUser struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

func (p *APIRouter) Auth(c *gin.Context) {
	var authUser AuthUser
	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.Error(common.NewError(http.StatusBadRequest, "JSON validation failed", err))
		return
	}

	user, err := p.UserService.FindUserByUsername(authUser.Username)
	if err != nil { // user is not exist
		c.Error(common.NewError(http.StatusForbidden, "Username or password is incorrect", err))
		return
	}

	if !user.CheckPassword(authUser.Password) {
		c.Error(common.NewError(http.StatusForbidden, "Username or password is incorrect", err))
		return
	}

	jwt, err := common.GetJWT(common.JWTUser{ID: user.ID, Username: user.Username})
	if err != nil {
		c.Error(common.NewError(http.StatusInternalServerError, "Something went wrong", err))
		return
	}

	response := AuthResponse{JWT: jwt}

	c.JSON(http.StatusOK, response)
}
