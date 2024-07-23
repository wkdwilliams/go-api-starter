package storage

import "go-api-starter/internal/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) GetUserById(id int) (types.User, error) {
	user := types.User{
		Firstname: "Lewis",
		Lastname:  "Williams",
	}

	return user, nil
}
