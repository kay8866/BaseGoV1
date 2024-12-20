package routes

import (
	"BaseGoV1/internal/controllers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterAPIRoutes(e *echo.Echo, userController *controllers.UserController) {
	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// User routes
	e.GET("/api/users", userController.GetUsers)
	e.POST("/api/users", userController.CreateUser)

	// Add more API routes as needed
}
