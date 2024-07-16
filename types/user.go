package types

type User struct {
	AbstractType
	Firstname 	string	`json:"first_name"`
	Lastname 	string	`json:"last_name"`
}