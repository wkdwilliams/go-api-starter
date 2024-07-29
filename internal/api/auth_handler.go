package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type authRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s ApiServer) loginHandler(c echo.Context) error {
	var authRequest authRequest

	if err := c.Bind(&authRequest); err != nil {
		return err
	}

	if err := s.echo.Validator.Validate(authRequest); err != nil {
		return err
	}

	user, err := s.service.GetUserByUsername(authRequest.Username)

	if err != nil {
		return echo.ErrUnauthorized
	}

	if user.Password != authRequest.Password {
		return echo.ErrUnauthorized
	}

	auth, err := s.auth.Authenticate(*user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		Token string
	}{
		Token: auth,
	})
}
