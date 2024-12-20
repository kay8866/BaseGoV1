package main

//go:generate go run github.com/swaggo/swag/cmd/swag init -g ./cmd/api/main.go -o ./docs

import (
	_ "BaseGoV1/docs" // This line is important
	"BaseGoV1/internal/controllers"
	"BaseGoV1/internal/models"
	"BaseGoV1/internal/repositories"
	"BaseGoV1/internal/routes"
	_ "github.com/gofiber/swagger" // swagger handler
	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// @title BaseGoV1 API
// @version 1.0
// @description This is a sample server for BaseGoV1.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

func main() {
	// MySQL connection string
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "basegouser:basegopassword@tcp(localhost:3306)/basegov1?charset=utf8mb4&parseTime=True&loc=Local"
	}

	// Initialize database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to auto migrate: ", err)
	}

	// Initialize Echo
	e := echo.New()

	// Swagger route
	e.GET("/swagger/*", swagger.WrapHandler)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize controllers
	userController := controllers.NewUserController(userRepo)

	// Register routes
	routes.RegisterWebRoutes(e)
	routes.RegisterAPIRoutes(e, userController)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
