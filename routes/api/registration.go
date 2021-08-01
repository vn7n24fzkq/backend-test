package api

import (
	"errors"
	"net/http"
	"regexp"
	"vn7n24fzkq/backend-test/common"
	"vn7n24fzkq/backend-test/dao"

	"github.com/gin-gonic/gin"
)

type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (p *RegisterUser) validate() *common.CommonError {
	usernameMatchFailMsg := "The length of username should between 4 and 20, and only contains alphabet character and number"
	passwordMatchFailMsg := "The length of password should between 8 and 20, and only contains alphabet character and number"
	if match, err := regexp.MatchString("^[a-zA-Z0-9]{4,8}$", p.Username); !match {
		if err == nil {
			err = errors.New(usernameMatchFailMsg)
		}
		return common.NewError(http.StatusBadRequest, usernameMatchFailMsg, err)
	}

	if match, err := regexp.MatchString("^[a-zA-Z0-9]{8,20}$", p.Password); !match {
		if err == nil {
			err = errors.New(passwordMatchFailMsg)
		}
		return common.NewError(http.StatusBadRequest, passwordMatchFailMsg, err)
	}

	return nil
}

func (p *APIRouter) Registration(c *gin.Context) {
	var registerUser RegisterUser
	if err := c.ShouldBindJSON(&registerUser); err != nil {
		c.Error(common.NewError(http.StatusBadRequest, "JSON validation failed", err))
		return
	}

	if commonErr := registerUser.validate(); commonErr != nil {
		c.Error(commonErr)
		return
	}

	if _, err := p.UserService.FindUserByUsername(registerUser.Username); err == nil {
		c.Error(common.NewError(
			http.StatusConflict,
			"Username had already been registered",
			errors.New("Username had already been registered"),
		))
		return
	}

	daoUser := dao.User{}
	daoUser.Username = registerUser.Username
	passwordDigest, err := daoUser.HashPassword(registerUser.Password)
	if err != nil {
		c.Error(common.NewError(http.StatusInternalServerError, "Something went wrong", err))
		return
	}
	daoUser.PasswordDigest = passwordDigest

	p.UserService.CreateUser(daoUser)
	if err != nil {
		c.Error(common.NewError(http.StatusInternalServerError, "Something went wrong", err))
		return
	}

	c.Status(http.StatusOK)
}
