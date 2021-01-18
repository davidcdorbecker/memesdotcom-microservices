package services

import (
	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	log "github.com/sirupsen/logrus"
	"memesdotcom-auth/domain"
	"memesdotcom-auth/infrastructure/repository/redis"
	"memesdotcom-auth/infrastructure/repository/users_api"
	"time"
)

type authService struct {
	usersAPI    users_api.UsersAPI
	redisClient redis.Redis
}

type AuthService interface {
	LoginWithUserCredentials(userCredentials *domain.AccessTokenRequest) (*domain.AccessToken, _errors.RestError)
}

func NewAuthService(usersAPI users_api.UsersAPI, redisClient redis.Redis) AuthService {
	return &authService{usersAPI, redisClient}
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

	expiration := accessToken.RefreshTokenExpirationTime - time.Now().UnixNano()
	if err := as.redisClient.Set(accessToken.RefreshToken, "1", time.Duration(expiration)); err != nil {
		return nil, err
	}

	return accessToken, nil
}
