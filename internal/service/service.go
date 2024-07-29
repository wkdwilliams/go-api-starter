package service

import (
	"go-api-starter/internal/storage"
)

type Service struct {
	db storage.Storage
}

func New() Service {
	return Service{
		db: storage.NewSqlStorage(),
	}
}