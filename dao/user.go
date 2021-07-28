package dao

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint      `gorm:"primarykey"`
	Username       string    `json:"username" binding:"required"`
	PasswordDigest string    `json:"-"`
	Salt           string    `json:"-"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Tasks          []Task    `json:"-" gorm:"foreignKey:UserID"`
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

type UserDAO struct {
	db *gorm.DB
}

func (p *UserDAO) CreateUser() {
	//TODO
}

func (p *UserDAO) GetUserByID(id int) (User, error) {
	var user User
	if err := p.db.Model(&user).First(id).Error; err != nil {
		return user, err
	}
	return user, nil

}

func (p *UserDAO) UpdateUser() {
	//TODO
}

func (p *UserDAO) DeleteUserByID() {
	//TODO
}
