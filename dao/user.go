package dao

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int       `gorm:"primarykey"`
	Username       string    `json:"username" binding:"required" gorm:"not null;size:128;uniqueIndex"`
	PasswordDigest string    `json:"-" gorm:"not null;size:128"`
	Salt           string    `json:"-" gorm:"not null;size:128"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Tasks          []Task    `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

type UserDAO struct {
	db *gorm.DB
}

func (p *UserDAO) CreateUser(user User) (User, error) {
	result := p.db.Create(&user)
	return user, result.Error
}

// FindOneUser(User{ID:1})
func (p *UserDAO) FindOneUser(condition User) (User, error) {
	var user User
	result := p.db.Where(condition).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (p User) Update(dao *UserDAO, user User) error {
	if p.ID != user.ID {
		return errors.New("You can not update th ID")
	}
	return dao.db.Save(&user).Error
}

func (p User) Delete(dao *UserDAO) error {
	return dao.db.Delete(p).Error
}
