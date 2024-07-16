package storage

import (
	"go-api-starter/pkg"
	"go-api-starter/types"
	"math"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqlStorage struct {
	db *gorm.DB
}

func NewSqlStorage() *SqlStorage {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return &SqlStorage{
		db: db,
	}
}

func (s *SqlStorage) GetDB() *gorm.DB {
	return s.db
}

func (s *SqlStorage) Paginate(value any, pagination *pkg.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func (s *SqlStorage) GetUserById(id int) (*types.User, error) {
	var user types.User

	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *SqlStorage) GetAllUsers() (*pkg.Pagination, error) {
	var users []types.User
	var pagination pkg.Pagination

	if err := s.db.Scopes(s.Paginate(users, &pagination, s.db)).Find(&users).Error; err != nil {
		return nil, err
	}

	pagination.Items = users

	return &pagination, nil
}

func (s *SqlStorage) CreateUser(user *types.User) error {
	return s.db.Save(user).Error
}
