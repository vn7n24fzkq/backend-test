package service

import "vn7n24fzkq/backend-test/dao"

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO: userDAO}

}

type UserService struct {
	userDAO *dao.UserDAO
}

// func AuthUser() (dao.User, error) {
//
// }

func (u *UserService) GetUserById(id int) (dao.User, error) {
	user, err := u.userDAO.GetUserByID(id)

	return user, err
}

// func CreateUser(id int) (dao.User, error) {
// 	user, err := dao.CreateUser(id)
//
// 	return user, err
// }
