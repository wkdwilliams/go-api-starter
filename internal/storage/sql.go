package storage

import (
	"fmt"
	"go-api-starter/internal/types"
	"go-api-starter/pkg/paginate"
	"log"
	"math"
	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlStorage struct {
	db *gorm.DB
}

func NewSqlStorage() SqlStorage {
	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	return SqlStorage{
		db: db,
	}
}

func (s SqlStorage) GetDB() *gorm.DB {
	return s.db
}

// property t must ba a type.
func (s SqlStorage) Count(t any) int64 {
	var count int64

	s.db.Model(t).Count(&count)

	return count
}

// Creates a record based on the given type.
// Property t must be a type.
func (s SqlStorage) Create(t any) error {
	return s.db.Save(t).Error
}

// Updates a record based on the given type
// Property t must be a type.
func (s SqlStorage) Update(t any) error {
	val := reflect.ValueOf(t)

	// Check if the object is a pointer and dereference it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Ensure we're dealing with a struct
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct but got %s", val.Kind())
	}

	// Get the field by name
	field := val.FieldByName("ID")

	// Check if the field exists
	if !field.IsValid() {
		return fmt.Errorf("no such field: ID")
	}

	// Check if id exists. If id does not exist, we return an error
	if err := s.db.First(t, field.Interface()).Error; err != nil {
		return err
	}

	return s.db.Save(t).Error
}

// Property t must be a type.
func (s SqlStorage) Delete(t any) error {
	return s.db.Delete(t).Error
}

// Gets a record by given id.
// Property t must be a pointer to a type
func (s SqlStorage) GetById(t any, id int) error {
	return s.db.First(t, id).Error
}

// Paginates a give gorm query & type
// Property t must be a type
func (s SqlStorage) Paginate(t any, pagination *paginate.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(t).Count(&totalRows)

	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func (s SqlStorage) GetUserById(id int) (*types.User, error) {
	var user types.User

	if err := s.GetById(&user, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s SqlStorage) GetUserByUsername(username string) (*types.User, error) {
	var user types.User

	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s SqlStorage) GetAllUsers(pagination *paginate.Pagination) error {
	var users []types.User

	if err := s.db.Scopes(s.Paginate(users, pagination, s.db)).Find(&users).Error; err != nil {
		return err
	}

	pagination.Items = users

	return nil
}

func (s SqlStorage) CreateUser(user *types.User) error {
	return s.Create(user)
}

func (s SqlStorage) UpdateUser(user *types.User) error {
	return s.Update(user)
}

func (s SqlStorage) GetTotalUserCount() int64 {
	return s.Count(&types.User{})
}

func (s SqlStorage) GetLastUser() (*types.User, error) {
	var user types.User

	if err := s.db.Last(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s SqlStorage) DeleteUser(user *types.User) error {
	return s.Delete(user)
}
