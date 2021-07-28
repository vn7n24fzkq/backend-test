package service

import "vn7n24fzkq/backend-test/dao"

func NewTaskService(taskDAO *dao.TaskDAO) *TaskService {
	return &TaskService{taskDAO: taskDAO}

}

type TaskService struct {
	taskDAO *dao.TaskDAO
}

// func (p *TaskService) CreateTask(task dao.Task) (dao.Task, error) {
//
// }
//
// func (p *TaskService) GetTaskByUserID(id int) (dao.Task, error) {
//
// }
//
// func (p *TaskService) UpdateTask(task dao.Task) error {
//
// }
//
// func (p *TaskService) DeleteTask(id int) error {
//
// }
