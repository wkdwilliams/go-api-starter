package api

import (
	"github.com/labstack/echo/v4"
)

func (s ApiServer) registerRoutes(e *echo.Echo) {
	e.POST("/login", s.loginHandler)
	
	userGroup := e.Group("/user", authMiddleware)
	userGroup.GET("/:id", s.getUserHandler)
	userGroup.GET("/", s.getUsersHandler)
	userGroup.POST("/", s.createUserHandler)
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}