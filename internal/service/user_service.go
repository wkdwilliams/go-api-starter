package service

import (
	"go-api-starter/internal/types"
	"go-api-starter/pkg/paginate"
)

func (s Service) GetUserById(id int) (*types.User, error) {
	return s.db.GetUserById(id)
}

func (s Service) GetUserByUsername(username string) (*types.User, error) {
	return s.db.GetUserByUsername(username)
}

func (s Service) GetAllUsers(pagination *paginate.Pagination) error {
	return s.db.GetAllUsers(pagination)
}

func (s Service) CreateUser(user *types.User) error {
	return s.db.CreateUser(user)
}

func (s Service) UpdateUser(user *types.User) error {
	return s.db.UpdateUser(user)
}

func (s Service) GetTotalUserCount() int64 {
	return s.db.GetTotalUserCount()
}

func (s Service) GetLastUser() (*types.User, error) {
	return s.db.GetLastUser()
}

func (s Service) DeleteUser(user *types.User) error {
	return s.db.DeleteUser(user)
}
