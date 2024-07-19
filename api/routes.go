package api

import "github.com/labstack/echo/v4"

func (s ApiServer) registerRoutes(e *echo.Echo) {
	e.GET("/:id", s.getUserHandler)
	e.GET("/", s.getUsersHandler)
	e.POST("/", s.createUserHandler)
}
