package api

import (
	"net/http"
	"strconv"
	"time"
	"vn7n24fzkq/backend-test/common"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" binding:"required"`
	CreatedAt time.Time `json:"createAt"`
}

func (p *APIRouter) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(common.NewError(http.StatusBadRequest, "URL query parameter parse error", err))
		return
	}

	user, err := p.UserService.GetUserByID(id)
	if err != nil || user.ID != p.GetCurrentUser(c).ID {
		c.Error(common.NewError(http.StatusForbidden, "Forbidden", err))
		return
	}

	c.JSON(http.StatusOK, UserResponse{ID: user.ID, Username: user.Username, CreatedAt: user.CreatedAt})
}
