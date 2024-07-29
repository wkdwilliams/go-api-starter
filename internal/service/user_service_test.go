package service

import (
	"go-api-starter/internal/types"
	"go-api-starter/pkg/paginate"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	service := New()

	// Get the number of users before we create one
	userBeforeCount := service.GetTotalUserCount()

	fname := gofakeit.FirstName()
	lname := gofakeit.LastName()

	user := types.User{
		Firstname: fname,
		Lastname:  lname,
		Username:  gofakeit.Username(),
		Password:  gofakeit.Password(true, false, false, false, false, 5),
	}

	assert.Nil(t, service.CreateUser(&user))

	// Get the number of users after we create
	userAfterCount := service.GetTotalUserCount()

	assert.Equal(t, userBeforeCount+1, userAfterCount)

	assert.Equal(t, user.Firstname, fname)
	assert.Equal(t, user.Lastname, lname)
}

func TestGetUserById(t *testing.T) {
	service := New()

	lastUser, err := service.GetLastUser()

	assert.Nil(t, err)

	user, err := service.GetUserById(int(lastUser.ID))

	assert.Nil(t, err)

	assert.Equal(t, user.ID, lastUser.ID)
}

func TestGetAllUsers(t *testing.T) {
	service := New()

	var paginated paginate.Pagination
	paginated.Page = 1

	err := service.GetAllUsers(&paginated)

	assert.Nil(t, err)

	totalUserCount := service.GetTotalUserCount()

	assert.Equal(t, int64(len(paginated.Items.([]types.User))), totalUserCount)
}

func TestDeleteUser(t *testing.T) {
	service := New()

	user, err := service.GetLastUser()

	assert.Nil(t, err)

	service.DeleteUser(user)

	_, err = service.GetUserById(int(user.ID))

	assert.NotNil(t, err)
}
