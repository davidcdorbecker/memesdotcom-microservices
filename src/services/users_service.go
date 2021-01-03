package services

import (
	"memesdotcom-users/domain"
	"memesdotcom-users/infrastructure/repository/db"
	"memesdotcom-users/utils/date_utils"
	_errors "memesdotcom-users/utils/error"
	"memesdotcom-users/utils/helpers"
)

type service struct {
	db db.DbRepository
}

type Service interface {
	CreateUser(user domain.User) _errors.RestError
}

func NewService(dbRepo db.DbRepository) Service {
	return &service{db: dbRepo}
}

func (s *service) CreateUser(user domain.User) _errors.RestError {
	//formating user
	user.Status = "active"
	user.DateCreated = date_utils.GetNowDbFormat()
	user.Password = helpers.Encrypt(user.Password)

	if err := s.db.CreateClient(user); err != nil {
		return err
	}

	return nil
}
