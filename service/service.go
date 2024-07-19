package service

import (
	"go-api-starter/pkg"
	"go-api-starter/storage"
	"go-api-starter/types"
)

type Service struct {
	db storage.Storage
}

func NewService() Service {
	return Service{
		db: storage.NewSqlStorage(),
	}
}

func (s Service) GetUserById(id int) (*types.User, error) {
	return s.db.GetUserById(id)
}

func (s Service) GetAllUsers() (*pkg.Pagination, error) {
	return s.db.GetAllUsers()
}

func (s Service) CreateUser(user *types.User) error {
	return s.db.CreateUser(user)
}
