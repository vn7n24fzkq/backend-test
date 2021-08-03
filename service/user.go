package service

import "vn7n24fzkq/backend-test/dao"

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{UserDAO: userDAO}
}

type UserService struct {
	UserDAO *dao.UserDAO
}


// should return the created user with id
func (p *UserService) CreateUser(user dao.User) (dao.User, error) {
	return p.UserDAO.CreateUser(user)
}

func (p *UserService) GetUserByID(id int) (dao.User, error) {
	user, err := p.UserDAO.FindOneUser(dao.User{ID: id})
	return user, err
}

func (p *UserService) FindUserByUsername(username string) (dao.User, error) {
	user, err := p.UserDAO.FindOneUser(dao.User{Username: username})
	return user, err
}

func (p *UserService) UpdateUserByID(id int, user dao.User) error {
	user.ID = id
	targetUser, err := p.GetUserByID(id)
	if err != nil {
		return err
	}
	return targetUser.Update(p.UserDAO, user)
}

func (p *UserService) DeleteUserByID(id int) error {
	user, err := p.GetUserByID(id)
	if err != nil {
		return err
	}
	return user.Delete(p.UserDAO)
}
