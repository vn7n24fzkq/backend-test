package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "something was wrong")
		return
	}
	c.JSON(200, map[string]int{"id": id})
}
