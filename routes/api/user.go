package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (p *APIRoutes) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "something was wrong")
		return
	}
	data, err := p.UserService.GetUserById(id)
	if err != nil {
		c.String(400, "something was wrong")
		return
	}
	c.JSON(200, data)
}

func (p *APIRoutes) CreateUser(c *gin.Context) {
	//TODO
}
