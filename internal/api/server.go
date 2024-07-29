package api

import (
	"context"
	"errors"
	"go-api-starter/internal/auth"
	"go-api-starter/internal/service"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ApiServer struct {
	listenAddr string
	echo       *echo.Echo
	service    service.Service
	auth       auth.Auth
}

func New(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		echo:       echo.New(),
		service:    service.New(),
		auth:       auth.NewJWT(),
	}
}

func (s ApiServer) Start() error {
	s.echo.HTTPErrorHandler = customHTTPErrorHandler
	s.echo.Validator = &customValidator{
		validator: *validator.New(validator.WithRequiredStructEnabled()),
	}

	// Register the routes in routes.go
	s.registerRoutes()

	return s.echo.Start(s.listenAddr)
}

func (s ApiServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.echo.Shutdown(ctx)
}

type ApiResponse struct {
	Data any `json:"data"`
}

func NewApiResponse(data any) ApiResponse {
	return ApiResponse{
		Data: data,
	}
}

type ValidationError struct {
	Message string   `json:"message"`
	Error   []string `json:"error"`
}

type customValidator struct {
	validator validator.Validate
}

func (cv customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, he)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Handle record not found error...
		c.JSON(http.StatusNotFound, echo.ErrNotFound)
		return
	} else if _, ok := err.(validator.ValidationErrors); ok {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, strings.Split(err.Error(), "Error:")[1])
		}
		c.JSON(http.StatusBadRequest, ValidationError{
			Message: "Validation error",
			Error:   errors,
		})
		return
	}

	// Default error is 500 - internal server error
	log.Err(err).Msg("")
	c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
}
