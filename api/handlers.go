package api

import (
	"go-api-starter/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getUserDTO struct {
	Id string `param:"id"`
}

func (s Server) getUserHandler(c echo.Context) error {
	var userDto getUserDTO

	if err := c.Bind(&userDto); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, userDto)
}

func (s Server) getUsersHandler(c echo.Context) error {
	users, err := s.db.GetAllUsers()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, users)
}

type createUserDTO struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
}

func (s Server) createUserHandler(c echo.Context) error {
	var userDTO createUserDTO

	if err := c.Bind(&userDTO); err != nil {
		return err
	}

	if err := s.validator.Struct(userDTO); err != nil {
		return err
	}

	userType := types.User{
		Firstname: userDTO.Firstname,
		Lastname:  userDTO.Lastname,
	}

	if err := s.db.CreateUser(&userType); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, userType)
}
