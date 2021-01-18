package handlers

import (
	"memesdotcom-auth/services"
	"net/http"

	"memesdotcom-auth/domain"

	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService services.AuthService
}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{authService}
}

func (ah *authHandler) Login(c *fiber.Ctx) error {
	var loginRequest domain.UserCredentials

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(_errors.NewBadRequestError("error in request body"))
	}

	if user, err := ah.authService.LoginWithEmailAndPassword(&loginRequest); err != nil {
		return c.Status(err.Code()).JSON(_errors.NewBadRequestError(err.Message()))
	} else {
		return c.Status(http.StatusOK).JSON(&user)
	}
}
