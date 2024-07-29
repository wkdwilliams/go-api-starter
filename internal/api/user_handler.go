package api

import (
	"go-api-starter/internal/types"
	"go-api-starter/pkg/paginate"
	"net/http"

	_ "go-api-starter/docs"

	"github.com/labstack/echo/v4"
)

// @Summary Get 	Get user by id
// @Description 	Get the user by given id
// @Tags 			User
// @Accept 			*/*
// @Produce 		json
// @Param 			id path int true "id"
// @Success 		200 {object} ApiResponse{data=types.User}
// @Failure 		500 {object} echo.HTTPError
// @Failure 		404 {object} echo.HTTPError
// @Failure 		400 {object} ValidationError
// @Router 			/user/{id} [get]
func (s ApiServer) getUserHandler(c echo.Context) error {
	var userDto types.GetUserDTO

	if err := c.Bind(&userDto); err != nil {
		return err
	}

	user, err := s.service.GetUserById(userDto.Id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, NewApiResponse(user))
}

// @Summary 		Get user by id
// @Description 	Get all users paginated
// @Tags 			User
// @Accept 			*/*
// @Produce 		json
// @Param			page query int false "page"
// @Success 		200 {object} paginate.Pagination{items=[]types.User}
// @Failure 		500 {object} echo.HTTPError
// @Failure 		403 {object} echo.HTTPError
// @Router 			/user [get]
func (s ApiServer) getUsersHandler(c echo.Context) error {
	var paginated paginate.Pagination

	if err := c.Bind(&paginated); err != nil {
		return err
	}

	if err := s.service.GetAllUsers(&paginated); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, paginated)
}

// @Summary 	Create user
// @Description Creates a user
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param 		request body types.CreateUserDTO true "body"
// @Success 	201 {object} ApiResponse{data=types.User}
// @Failure 	500 {object} echo.HTTPError
// @Failure 	400 {object} ValidationError
// @Router 		/user [post]
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
		Username:  userDTO.Username,
		Password:  userDTO.Password,
	}

	if err := s.service.CreateUser(&userType); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, NewApiResponse(userType))
}

// @Summary 	Update user
// @Description Update a user
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param 		request body types.UpdateUserDTO true "body"
// @Success 	200 {object} ApiResponse{data=types.User}
// @Failure 	500 {object} echo.HTTPError
// @Failure 	404 {object} echo.HTTPError
// @Failure 	400 {object} ValidationError
// @Router 		/user [put]
func (s ApiServer) updateUserHandler(c echo.Context) error {
	var userDTO types.UpdateUserDTO

	if err := c.Bind(&userDTO); err != nil {
		return err
	}

	if err := c.Validate(userDTO); err != nil {
		return err
	}

	userType := types.User{
		AbstractType: types.AbstractType{
			ID: userDTO.Id,
		},
		Firstname: userDTO.Firstname,
		Lastname:  userDTO.Lastname,
		Username:  userDTO.Username,
		Password:  userDTO.Password,
	}

	if err := s.service.UpdateUser(&userType); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userType)
}
