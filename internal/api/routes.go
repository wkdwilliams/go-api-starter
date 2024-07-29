package api

import (
	_ "go-api-starter/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s ApiServer) registerRoutes() {

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	s.echo.POST("/login", s.loginHandler)

	userGroup := s.echo.Group("/user", authMiddleware)
	userGroup.GET("/:id", s.getUserHandler)
	userGroup.GET("", s.getUsersHandler)
	userGroup.PUT("", s.updateUserHandler)
	userGroup.POST("", s.createUserHandler)
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
