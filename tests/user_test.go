package tests

import (
	"go-api-starter/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	user, err := service.GetUserById(1)

	assert.Nil(t, err)

	assert.Equal(t, user.ID, uint(1))
}

func TestGetAllUsers(t *testing.T) {
	users, err := service.GetAllUsers()

	assert.Nil(t, err)

	user := users.Items.([]types.User)[0]

	assert.Equal(t, user.ID, uint(1))
}

func TestCreateUser(t *testing.T) {
	user := types.User{
		Firstname: "test",
		Lastname:  "user",
	}

	assert.Nil(t, service.CreateUser(&user))

	assert.Equal(t, user.Firstname, "test")
	assert.Equal(t, user.Lastname, "user")
}
