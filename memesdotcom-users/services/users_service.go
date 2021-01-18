package services

import (
	"memesdotcom-users/domain"
	"memesdotcom-users/infrastructure/repository/db"
	"memesdotcom-users/utils/constants"
	"memesdotcom-users/utils/date_utils"
	"memesdotcom-users/utils/helpers"

	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
)

type service struct {
	db db.DbRepository
}

type Service interface {
	CreateUser(user domain.User) _errors.RestError
	VerifyUserCredentials(user *domain.UserCredentials) (*domain.User, _errors.RestError)
}

func NewService(dbRepo db.DbRepository) Service {
	return &service{db: dbRepo}
}

func (s *service) CreateUser(user domain.User) _errors.RestError {
	//formating user
	user.Status = constants.StatusActive
	user.DateCreated = date_utils.GetNowDbFormat()
	user.Password = helpers.Encrypt(user.Password)

	if err := s.db.CreateClient(user); err != nil {
		return err
	}

	return nil
}

func (s *service) VerifyUserCredentials(userCredentials *domain.UserCredentials) (*domain.User, _errors.RestError) {
	userCredentials.Password = helpers.Encrypt(userCredentials.Password)

	if user, err := s.db.FindByCredentials(userCredentials); err != nil {
		return nil, err
	} else {
		return user, nil
	}

}
