package types

// User model info
// @Description User account information
type User struct {
	AbstractType
	Firstname string `json:"firstname" gorm:"not null"`
	Lastname  string `json:"lastname" gorm:"not null"`
	Username  string `json:"-" gorm:"not null"`
	Password  string `json:"-" gorm:"not null"`
}

type CreateUserDTO struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type UpdateUserDTO struct {
	Id        uint   `json:"id" validate:"number"`
	Firstname string `json:"firstname" validate:""`
	Lastname  string `json:"lastname" validate:""`
	Username  string `json:"username" validate:""`
	Password  string `json:"password" validate:""`
}

type GetUserDTO struct {
	Id int `param:"id"`
}
