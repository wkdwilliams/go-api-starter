package storage

import (
	"go-api-starter/pkg"
	"go-api-starter/internal/types"
)

type Storage interface {
	GetUserById(int) (*types.User, error)
	GetAllUsers() (*pkg.Pagination, error)
	CreateUser(*types.User) error
	GetUserByUsername(string) (*types.User, error)
}
