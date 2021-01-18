package domain

import (
	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	"github.com/form3tech-oss/jwt-go"
	"github.com/spf13/viper"
	"memesdotcom-auth/utils/constants"
	"time"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AccessToken struct {
	UserID                     int64  `json:"user_id,omitempty"`
	ClientID                   string `json:"client_id,omitempty"`
	AccessToken                string `json:"access_token"`
	RefreshToken               string `json:"refresh_token"`
	AccessTokenExpirationTime  int64  `json:"access_token_expiration_time"`
	RefreshTokenExpirationTime int64  `json:"refresh_token_expiration_time"`
}

func (at *AccessToken) Generate() _errors.RestError {
	accessTokenExpirationTime := time.Now().Add(time.Minute * viper.GetDuration(constants.AccessTokenExpirationTime)).Unix()
	refreshTokenExpirationTime := time.Now().Add(time.Hour * viper.GetDuration(constants.RefreshTokenExpirationTime)).Unix()

	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	accessTokenClaims["userID"] = at.UserID
	accessTokenClaims["clientID"] = at.ClientID
	accessTokenClaims["exp"] = accessTokenExpirationTime

	aToken, atErr := accessToken.SignedString([]byte(viper.GetString(constants.AccessTokenSecret)))
	if atErr != nil {
		return _errors.NewInternalServerError("access token generation error")
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := accessToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["exp"] = refreshTokenExpirationTime
	rToken, rtErr := refreshToken.SignedString([]byte(viper.GetString(constants.RefreshTokenSecret)))
	if rtErr != nil {
		return _errors.NewInternalServerError("refresh token generation error")
	}

	at.AccessToken = aToken
	at.AccessTokenExpirationTime = accessTokenExpirationTime
	at.RefreshToken = rToken
	at.RefreshTokenExpirationTime = refreshTokenExpirationTime

	return nil
}
