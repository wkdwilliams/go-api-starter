package types

type User struct {
	AbstractType
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

type CreateUserDTO struct {
	Firstname string `json:"first_name" validate:"required"`
	Lastname  string `json:"last_name" validate:"required"`
}

type GetUserDTO struct {
	Id   int    `param:"id"`
}
