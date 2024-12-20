package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterWebRoutes(e *echo.Echo) {
	// Add web routes here
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to BaseGoV1")
	})

	// Add more web routes as needed
}
