package dao

import (
	"time"
)

type TodoList struct {
	ID        uint      `gorm:"primarykey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ExpiredAt time.Time `json:"expiredAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uint      `json:"UserID"`
}
