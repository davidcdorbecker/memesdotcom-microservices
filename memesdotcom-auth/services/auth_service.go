package services

import (
	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	"memesdotcom-auth/domain"
	"memesdotcom-auth/infrastructure/repository/users_api"
)

type authService struct {
	usersAPI users_api.UsersAPI
}

type AuthService interface {
	LoginWithEmailAndPassword(userCredentials *domain.UserCredentials) (*domain.User, _errors.RestError)
}

func NewAuthService(usersAPI users_api.UsersAPI) AuthService {
	return &authService{usersAPI}
}

func (as *authService) LoginWithEmailAndPassword(userCredentials *domain.UserCredentials) (*domain.User, _errors.RestError) {
	return as.usersAPI.GetUser(userCredentials)
}
