package api

import (
	"errors"
	"net/http"
	"strconv"
	"vn7n24fzkq/backend-test/common"
	"vn7n24fzkq/backend-test/dao"

	"github.com/gin-gonic/gin"
)

type TaskRequest struct {
	Title     string `json:"title" binding:"required,max=20"`
	Content   string `json:"content" binding:"required,max=200"`
	ExpiredAt int64  `json:"expiredAt"`
	Done      bool   `json:"done"`
}

func (p *APIRouter) GetAllTasks(c *gin.Context) {
	currentUser := p.GetCurrentUser(c)
	tasks, err := p.TaskService.GetAllTaskByUserID(currentUser.ID)
	if err != nil {
		c.Error(common.NewError(http.StatusInternalServerError, "Something went wrong", err))
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (p *APIRouter) CreateTask(c *gin.Context) {
	var task TaskRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		c.Error(common.NewError(http.StatusBadRequest, "JSON validation failed", err))
		return
	}

	currentUser := p.GetCurrentUser(c)

	daoTask, err := p.TaskService.CreateTask(
		dao.Task{
			Title:     task.Title,
			Content:   task.Content,
			ExpiredAt: task.ExpiredAt,
			Done:      task.Done,
			UserID:    currentUser.ID,
		})

	if err != nil {
		c.Error(common.NewError(http.StatusInternalServerError, "Something went wrong", err))
		return
	}
	c.JSON(http.StatusOK, daoTask)
}

func (p *APIRouter) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(common.NewError(http.StatusBadRequest, "URL query parameter parse error", err))
		return
	}
	var taskRequest TaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.Error(common.NewError(http.StatusBadRequest, "JSON validation failed", err))
		return
	}

	targetTask, taskErr := p.TaskService.GetTaskByID(id)
	if taskErr != nil {
		c.Error(common.NewError(http.StatusNotFound, "Task is not exist", taskErr))
		return
	}
	currentUser := p.GetCurrentUser(c)
	if currentUser.ID != targetTask.UserID {
		c.Error(common.NewError(http.StatusForbidden, "Forbidden", errors.New("Forbidden")))
		return
	}
	targetTask.Title = taskRequest.Title
	targetTask.Content = taskRequest.Content
	targetTask.ExpiredAt = taskRequest.ExpiredAt
	targetTask.Done = taskRequest.Done
	p.TaskService.UpdateTaskByID(id, targetTask)
}

func (p *APIRouter) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(common.NewError(http.StatusBadRequest, "URL query parameter parse error", err))
		return
	}

	task, taskErr := p.TaskService.GetTaskByID(id)
	if taskErr != nil {
		c.Error(common.NewError(http.StatusNotFound, "Task is not exist", taskErr))
		return
	}
	currentUser := p.GetCurrentUser(c)
	if currentUser.ID != task.UserID {
		c.Error(common.NewError(http.StatusForbidden, "Forbidden", errors.New("Forbidden")))
		return
	}

	deletedErr := p.TaskService.DeleteTaskByID(id)
	if deletedErr != nil {
		c.Error(common.NewError(http.StatusInternalServerError, "Something went wrong", deletedErr))
		return
	}

	c.Status(http.StatusOK)
}
