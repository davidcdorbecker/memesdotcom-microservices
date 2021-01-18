package services

import (
	"fmt"
	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	"github.com/form3tech-oss/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"memesdotcom-auth/domain"
	"memesdotcom-auth/infrastructure/repository/redis"
	"memesdotcom-auth/infrastructure/repository/users_api"
	"memesdotcom-auth/utils/constants"
	"strconv"
)

const (
	accessTokenError = "access token error"
)

type authService struct {
	usersAPI    users_api.UsersAPI
	redisClient redis.Redis
}

type AuthService interface {
	LoginWithUserCredentials(userCredentials *domain.AccessTokenRequest) (*domain.AccessToken, _errors.RestError)
	ValidateAndGenerateAccessToken(accessToken *domain.AccessToken) (*domain.AccessToken, _errors.RestError)
	GenerateAccessToken(accessToken *domain.AccessToken) (*domain.AccessToken, _errors.RestError)
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

	accessToken, err := as.createAccessToken(user.ID, "")
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (as *authService) ValidateAndGenerateAccessToken(accessToken *domain.AccessToken) (*domain.AccessToken, _errors.RestError) {

	at, err := jwt.Parse(accessToken.AccessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
			return nil, _errors.NewBadRequestError(accessTokenError)
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(viper.GetString(constants.AccessTokenSecret)), nil
	})
	if err != nil || at == nil {
		return nil, _errors.NewBadRequestError(accessTokenError)
	}

	if claims, ok := at.Claims.(jwt.MapClaims); ok && at.Valid {
		accessToken.RefreshToken = claims["refreshToken"].(string)
		userID, _ := strconv.ParseInt(claims["userID"].(string), 10, 64)
		accessToken.UserID = userID
		accessToken.ClientID = claims["clientID"].(string)
		expiration, _ := strconv.ParseInt(claims["exp"].(string), 10, 64)
		accessToken.AccessTokenExpirationTime = expiration

		if accessToken.IsExpired() || !as.isRefreshTokenAvailable(accessToken) {
			return nil, _errors.NewNotFoundError("access token expired")
		}

	} else {
		return nil, _errors.NewBadRequestError(accessTokenError)
	}

	newAccessToken, restErr := as.generateNewAccessToken(accessToken)
	if restErr != nil {
		return nil, restErr
	}

	return newAccessToken, nil
}

func (as *authService) GenerateAccessToken(accessToken *domain.AccessToken) (*domain.AccessToken, _errors.RestError) {

	at, err := jwt.Parse(accessToken.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
			return nil, _errors.NewBadRequestError(accessTokenError)
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(viper.GetString(constants.RefreshTokenSecret)), nil
	})
	if err != nil || at == nil {
		return nil, _errors.NewBadRequestError(accessTokenError)
	}

	if claims, ok := at.Claims.(jwt.MapClaims); ok && at.Valid {
		userID, _ := strconv.ParseInt(claims["userID"].(string), 10, 64)
		accessToken.UserID = userID
		accessToken.ClientID = claims["clientID"].(string)
		expiration, _ := strconv.ParseInt(claims["exp"].(string), 10, 64)
		accessToken.RefreshTokenExpirationTime = expiration

		if !as.isRefreshTokenAvailable(accessToken) || accessToken.IsRefreshTokenExpired() {
			return nil, _errors.NewNotFoundError("access token expired")
		}

	} else {
		return nil, _errors.NewBadRequestError(accessTokenError)
	}

	newAccessToken, restErr := as.generateNewAccessToken(accessToken)
	if restErr != nil {
		return nil, restErr
	}

	return newAccessToken, nil
}

func (as *authService) createAccessToken(userID int64, clientID string) (*domain.AccessToken, _errors.RestError) {
	accessToken := &domain.AccessToken{
		UserID:   userID,
		ClientID: clientID,
	}
	if err := accessToken.Generate(); err != nil {
		log.Error(err)
		return nil, _errors.NewInternalServerError("login failure")
	}
	accessToken.UserID = 0
	accessToken.ClientID = ""

	if err := as.redisClient.Set(accessToken.RefreshToken, "1", accessToken.GetExpiration()); err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (as *authService) generateNewAccessToken(accessToken *domain.AccessToken) (*domain.AccessToken, _errors.RestError) {
	if err := as.redisClient.Delete(accessToken.RefreshToken); err != nil {
		return nil, err
	}

	if err := accessToken.Generate(); err != nil {
		return nil, err
	}

	if err := as.redisClient.Set(accessToken.RefreshToken, "1", accessToken.GetExpiration()); err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (as *authService) isRefreshTokenAvailable(accessToken *domain.AccessToken) bool {
	rt, err := as.redisClient.Get(accessToken.RefreshToken)
	if err != nil || rt != "1" {
		return false
	}
	return true
}
