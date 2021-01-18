package domain

import (
	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	"github.com/form3tech-oss/jwt-go"
	"github.com/spf13/viper"
	"memesdotcom-auth/utils/constants"
	"strconv"
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
	accessTokenExpirationTime := time.Now().Add(time.Second * viper.GetDuration(constants.AccessTokenExpirationTime)).UnixNano()
	refreshTokenExpirationTime := time.Now().Add(time.Hour * viper.GetDuration(constants.RefreshTokenExpirationTime)).UnixNano()

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["exp"] = refreshTokenExpirationTime
	rToken, rtErr := refreshToken.SignedString([]byte(viper.GetString(constants.RefreshTokenSecret)))
	if rtErr != nil {
		return _errors.NewInternalServerError("refresh token generation error")
	}

	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	uid := strconv.FormatInt(at.UserID, 10)
	accessTokenClaims["userID"] = uid
	accessTokenClaims["clientID"] = at.ClientID
	exp := strconv.FormatInt(accessTokenExpirationTime, 10)
	accessTokenClaims["exp"] = exp
	accessTokenClaims["refreshToken"] = rToken

	aToken, atErr := accessToken.SignedString([]byte(viper.GetString(constants.AccessTokenSecret)))
	if atErr != nil {
		return _errors.NewInternalServerError("access token generation error")
	}

	at.AccessToken = aToken
	at.AccessTokenExpirationTime = accessTokenExpirationTime
	at.RefreshToken = rToken
	at.RefreshTokenExpirationTime = refreshTokenExpirationTime

	return nil
}

func (at *AccessToken) IsExpired() bool {
	return at.AccessTokenExpirationTime-time.Now().UnixNano() < 0
}

func (at *AccessToken) GetExpiration() time.Duration {
	return time.Duration(at.RefreshTokenExpirationTime - time.Now().UnixNano())
}
