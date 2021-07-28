package dao

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	ExpiredAt time.Time `json:"expiredAt" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uint      `json:"UserID"`
}

type TaskDAO struct {
	db *gorm.DB
}

func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{db: db}
}
