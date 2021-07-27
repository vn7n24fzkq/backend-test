package api

import (
	"strconv"
	"vn7n24fzkq/backend-test/service"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "something was wrong")
		return
	}
	data, err := service.GetUserById(id)
	if err != nil {
		c.String(400, "something was wrong")
		return
	}
	c.JSON(200, data)
}

func CreateUser(c *gin.Context) {
	//TODO
}
