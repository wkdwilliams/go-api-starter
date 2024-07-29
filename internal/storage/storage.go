package storage

import (
	"go-api-starter/internal/types"
	"go-api-starter/pkg/paginate"
)

type Storage interface {
	GetUserById(int) (*types.User, error)
	GetAllUsers(*paginate.Pagination) error
	CreateUser(*types.User) error
	UpdateUser(*types.User) error
	DeleteUser(*types.User) error
	GetUserByUsername(string) (*types.User, error)
	GetTotalUserCount() int64
	GetLastUser() (*types.User, error)
}
