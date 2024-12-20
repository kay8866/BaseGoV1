package controllers

import (
	"BaseGoV1/internal/models"
	"BaseGoV1/internal/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	UserRepo *repositories.UserRepository
}

func NewUserController(userRepo *repositories.UserRepository) *UserController {
	return &UserController{UserRepo: userRepo}
}

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of results"
// @Success 200 {array} models.User
// @Router /users [get]
func (uc *UserController) GetUsers(c echo.Context) error {
	users, err := uc.UserRepo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}
	return c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Create user"
// @Success 201 {object} models.User
// @Router /users [post]
func (uc *UserController) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := uc.UserRepo.Create(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, user)
}
