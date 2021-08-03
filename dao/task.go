package dao

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID         int       `gorm:"primarykey"`
	Title      string    `gorm:"not null" json:"title"`
	Content    string    `gorm:"not null" json:"content"`
	ExpiredAt  int64     `gorm:"index" json:"expiredAt"`
	Done       bool      `json:"done"`
	SendNotify bool      `gorm:"not null;default:false" json:"-"`
	CreatedAt  time.Time `json:"createAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	UserID     int       `gorm:"index;not null" json:"userID"`
}

type TaskDAO struct {
	db *gorm.DB
}

func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{db: db}
}

func (p *TaskDAO) CreateTask(task Task) (Task, error) {
	result := p.db.Debug().Create(&task)
	return task, result.Error
}

// FindOneTask(Task{ID:1})
func (p *TaskDAO) FindOneTask(condition Task) (Task, error) {
	var task Task
	result := p.db.Where(condition).First(&task)
	if result.Error != nil {
		return task, result.Error
	}

	return task, nil
}

// FindTasks(Task{UserID:1})
func (p *TaskDAO) FindTasks(condition Task) ([]Task, error) {
	var tasks []Task
	result := p.db.Order("done").Order("expired_at").Where(condition).Find(&tasks)
	if result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}

func (p *TaskDAO) FindNeedNotifyTasks(expiredUnixTimestamp int64) ([]Task, error) {
	var tasks []Task
	result := p.db.Order("expired_at").Where("done = ? AND send_notify = ? AND expired_at < ? ",
		false, false, expiredUnixTimestamp).Find(&tasks)
	if result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}

func (p Task) Update(dao *TaskDAO, task Task) error {
	if p.ID != task.ID {
		return errors.New("You can not update th ID")
	}
	return dao.db.Save(&task).Error
}

func (p Task) Delete(dao *TaskDAO) error {
	return dao.db.Delete(p).Error
}
