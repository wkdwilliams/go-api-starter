package api

import (
	"context"
	"errors"
	"fmt"
	"go-api-starter/storage"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	listenAddr string
	Echo       *echo.Echo
	db         storage.Storage
	validator  *validator.Validate
}

func New(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		Echo:       echo.New(),
		db:         storage.NewSqlStorage(),
		validator:  validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (s Server) Start() error {
	s.Echo.HTTPErrorHandler = customHTTPErrorHandler
	s.Echo.GET("/:id", s.getUserHandler)
	s.Echo.GET("/", s.getUsersHandler)
	s.Echo.POST("/", s.createUserHandler)

	return s.Echo.Start(s.listenAddr)
}

func (s Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.Echo.Shutdown(ctx)
}

func NewNotFoundError() *echo.HTTPError {
	return &echo.HTTPError{Code: 404, Message: http.StatusText(404)}
}

type ValidationError struct {
	Message string   `json:"message"`
	Error   []string `json:"error"`
}

func customHTTPErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, he)
		return
	} else {
		// Here we define our custom errors
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle record not found error...
			c.JSON(http.StatusNotFound, echo.ErrNotFound)
			return
		} else if _, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, fmt.Sprintf("Field %s must be %s", err.Field(), err.Tag()))
			}
			c.JSON(http.StatusBadRequest, ValidationError{
				Message: "Validation error",
				Error:   errors,
			})
			return
		}
	}

	// Default error is 500 - internal server error
	log.Println(err.Error())
	c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
}
