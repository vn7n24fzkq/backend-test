package service

import "vn7n24fzkq/backend-test/dao"

func GetUserById(id int) (dao.User, error) {
	user, err := dao.GetUserByID(id)

	return user, err
}
