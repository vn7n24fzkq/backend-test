package service

import (
	"vn7n24fzkq/backend-test/dao"
)

func NewTaskService(taskDAO *dao.TaskDAO, userService *UserService) *TaskService {
	return &TaskService{taskDAO: taskDAO, userService: userService}
}

type TaskService struct {
	taskDAO     *dao.TaskDAO
	userService *UserService
}

func (p *TaskService) CreateTask(task dao.Task) (dao.Task, error) {
	_, err := p.userService.GetUserByID(task.UserID)
	if err != nil {
		return task, err
	}
	return p.taskDAO.CreateTask(task)
}

func (p *TaskService) GetTaskByID(id int) (dao.Task, error) {
	return p.taskDAO.FindOneTask(dao.Task{ID: id})
}

func (p *TaskService) GetAllTaskByUserID(id int) ([]dao.Task, error) {
	return p.taskDAO.FindTasks(dao.Task{UserID: id})
}

func (p *TaskService) UpdateTaskByID(id int, task dao.Task) error {
	task.ID = id
	targetTask, err := p.GetTaskByID(id)
	if err != nil {
		return err
	}
	return targetTask.Update(p.taskDAO, task)
}

func (p *TaskService) DeleteTaskByID(id int) error {
	task, err := p.GetTaskByID(id)
	if err != nil {
		return err
	}
	return task.Delete(p.taskDAO)
}
