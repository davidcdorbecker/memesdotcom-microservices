package handlers

import (
	"memesdotcom-auth/services"
	"net/http"

	"memesdotcom-auth/domain"

	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	"github.com/gofiber/fiber/v2"
)

const (
	loginWithUserCredentials = "LOGIN_USER_CREDENTIALS"
	loginWithClientID        = "LOGIN_CLIENT-ID_CLIENT-SECRET"
	loginNotSupported        = "LOGIN_NOT_SUPPORTED"
)

type authHandler struct {
	authService services.AuthService
}

type AuthHandler interface {
	CreateAccessToken(c *fiber.Ctx) error
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{authService}
}

func (ah *authHandler) CreateAccessToken(c *fiber.Ctx) error {
	var loginRequest domain.AccessTokenRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(_errors.NewBadRequestError("error in request body"))
	}

	method := getLoginMethod(&loginRequest)
	switch method {
	case loginWithUserCredentials:
		if user, err := ah.authService.LoginWithUserCredentials(&loginRequest); err != nil {
			return c.Status(err.Code()).JSON(_errors.NewBadRequestError(err.Message()))
		} else {
			return c.Status(http.StatusOK).JSON(&user)
		}
	}

	restError := _errors.NewBadRequestError("method not supported")
	return c.Status(restError.Code()).JSON(restError.Message())
}

func getLoginMethod(accessTokenRequest *domain.AccessTokenRequest) string {
	if (accessTokenRequest.Username != "" || accessTokenRequest.Email != "") && accessTokenRequest.Password != "" {
		return loginWithUserCredentials
	} else if accessTokenRequest.ClientId != "" && accessTokenRequest.ClientSecret != "" {
		return loginWithClientID
	} else {
		return loginNotSupported
	}

}
