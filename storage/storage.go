package storage

import (
	"go-api-starter/pkg"
	"go-api-starter/types"
)

type Storage interface {
	GetUserById(int) (*types.User, error)
	GetAllUsers() (*pkg.Pagination, error)
	CreateUser(*types.User) error
}
