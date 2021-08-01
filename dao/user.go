package dao

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID             int    `gorm:"primarykey"`
	Username       string `gorm:"not null;size:20;uniqueIndex"`
	PasswordDigest string `gorm:"not null;size:60"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Tasks          []Task `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

func (p *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.PasswordDigest), []byte(password))
	return err == nil
}

func (p *User) HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(bytes), err
}
