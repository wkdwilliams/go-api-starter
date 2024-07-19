package api

import (
	"go-api-starter/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s ApiServer) getUserHandler(c echo.Context) error {
	var userDto types.GetUserDTO

	if err := c.Bind(&userDto); err != nil {
		return err
	}

	user, err := s.service.GetUserById(userDto.Id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (s ApiServer) getUsersHandler(c echo.Context) error {
	users, err := s.service.GetAllUsers()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (s ApiServer) createUserHandler(c echo.Context) error {
	var userDTO types.CreateUserDTO

	if err := c.Bind(&userDTO); err != nil {
		return err
	}

	if err := c.Validate(userDTO); err != nil {
		return err
	}

	userType := types.User{
		Firstname: userDTO.Firstname,
		Lastname:  userDTO.Lastname,
	}

	if err := s.service.CreateUser(&userType); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, userType)
}
