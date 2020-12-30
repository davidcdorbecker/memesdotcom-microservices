package handlers

import (
	"net/http"

	"memesdotcom-users/domain"
	"memesdotcom-users/services"
	_errors "memesdotcom-users/utils/error"

	"github.com/gofiber/fiber/v2"
)

type usersHandler struct {
	service services.Service
}

type UsersHandler interface {
	CreateUser(c *fiber.Ctx) error
}

func NewUsersHandler(service services.Service) UsersHandler {
	return &usersHandler{service}
}

func (uh *usersHandler) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(_errors.NewBadRequestError("error in request body"))
	}
	if user.Email == "" {
		return c.Status(http.StatusBadRequest).JSON(_errors.NewBadRequestError("user required"))
	}
	if user.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(_errors.NewBadRequestError("password required"))
	}
	if user.FirstName == "" {
		return c.Status(http.StatusBadRequest).JSON(_errors.NewBadRequestError("firstname required"))
	}
	if user.LastName == "" {
		return c.Status(http.StatusBadRequest).JSON(_errors.NewBadRequestError("lastname required"))
	}

	if err := uh.service.CreateUser(user); err != nil {
		return c.JSON(err)
	}

	return c.JSON(user)
}
