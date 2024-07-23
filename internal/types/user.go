package types

type User struct {
	AbstractType
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Username  string `json:"-"`
	Password  string `json:"-"`
}

type CreateUserDTO struct {
	Firstname string `json:"first_name" validate:"required"`
	Lastname  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type GetUserDTO struct {
	Id int `param:"id"`
}
