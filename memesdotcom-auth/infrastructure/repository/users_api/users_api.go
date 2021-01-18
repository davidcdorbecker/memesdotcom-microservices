package users_api

import (
	"encoding/json"
	"fmt"
	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	"github.com/go-resty/resty/v2"
	"memesdotcom-auth/domain"
	"memesdotcom-auth/utils/constants"
)

type usersAPI struct {
	restClient *resty.Client
}

type UsersAPI interface {
	GetUserByCredentials(userCredentials *domain.UserLoginByCredentials) (*domain.User, _errors.RestError)
}

func NewUsersAPI(restClient *resty.Client) UsersAPI {
	return &usersAPI{restClient}
}

func (uAPI *usersAPI) GetUserByCredentials(userCredentials *domain.UserLoginByCredentials) (*domain.User, _errors.RestError) {
	var user domain.User

	req, err := json.Marshal(&userCredentials)
	if err != nil {
		return nil, _errors.NewInternalServerError("error when trying to marshal request")
	}

	fmt.Println(string(req))
	res, err := uAPI.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(req)).
		SetResult(&user).
		Post(constants.VerifyUserEndpoint)

	fmt.Println(res)

	if err != nil || !res.IsSuccess() {
		return nil, _errors.NewNotFoundError("user not found")
	}

	return &user, nil
}
