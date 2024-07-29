package storage

import (
	"errors"
	"go-api-starter/internal/types"
	"go-api-starter/pkg/paginate"
	"math"
)

type MemoryStorage struct {
	users []types.User
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s MemoryStorage) Paginate(pagination *paginate.Pagination) {
	pagination.TotalRows = int64(len(s.users))
	pagination.TotalPages = int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.GetLimit())))

}

func (s MemoryStorage) GetUserById(id int) (*types.User, error) {
	for _, v := range s.users {
		if v.ID == uint(id) {
			return &v, nil
		}
	}

	return nil, errors.New("not found")
}

func (s MemoryStorage) GetUserByUsername(username string) (*types.User, error) {
	for _, v := range s.users {
		if v.Username == username {
			return &v, nil
		}
	}

	return nil, errors.New("not found")
}

func (s MemoryStorage) GetAllUsers() (*paginate.Pagination, error) {
	var pagination paginate.Pagination

	s.Paginate(&pagination)

	pagination.Items = s.users

	return &pagination, nil
}

func (s *MemoryStorage) CreateUser(user *types.User) error {
	s.users = append(s.users, *user)

	return nil
}

func (s *MemoryStorage) UpdateUser(user *types.User) error {
	for i, v := range s.users {
		if v.ID == user.ID {
			s.users[i] = *user
			return nil
		}
	}

	return errors.New("not found")
}
