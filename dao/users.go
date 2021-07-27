package dao

import "time"

type User struct {
	ID             uint       `gorm:"primarykey"`
	Username       string     `json:"username"`
	PasswordDigest string     `json:"-"`
	Salt           string     `json:"-"`
	ExpiredAt      time.Time  `json:"expiredAt"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	TodoLists      []TodoList `json:"-" gorm:"foreignKey:UserID"`
}

func CreateUser() {
	//TODO
}

func GetUserByID(id int) (User, error) {
	var user User
	if err := db.Model(&user).First(id).Error; err != nil {
		return user, err
	}
	return user, nil

}

func UpdateUser() {
	//TODO
}

func DeleteUser() {
	//TODO
}