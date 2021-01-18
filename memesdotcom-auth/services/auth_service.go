package services

import (
	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	log "github.com/sirupsen/logrus"
	"memesdotcom-auth/domain"
	"memesdotcom-auth/infrastructure/repository/users_api"
)

type authService struct {
	usersAPI users_api.UsersAPI
}

type AuthService interface {
	LoginWithUserCredentials(userCredentials *domain.AccessTokenRequest) (*domain.AccessToken, _errors.RestError)
}

func NewAuthService(usersAPI users_api.UsersAPI) AuthService {
	return &authService{usersAPI}
}

func (as *authService) LoginWithUserCredentials(accessTokenRequest *domain.AccessTokenRequest) (*domain.AccessToken, _errors.RestError) {
	userCredentials := &domain.UserLoginByCredentials{
		Username: accessTokenRequest.Username,
		Email:    accessTokenRequest.Email,
		Password: accessTokenRequest.Password,
	}

	user, err := as.usersAPI.GetUserByCredentials(userCredentials)
	if err != nil {
		return nil, err
	}

	accessToken := &domain.AccessToken{
		UserID: user.ID,
	}
	if err := accessToken.Generate(); err != nil {
		log.Error(err)
		return nil, _errors.NewInternalServerError("login failure")
	}
	accessToken.UserID = 0
	accessToken.ClientID = ""

	return accessToken, nil
}
