package db

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"memesdotcom-users/domain"
	_errors "memesdotcom-users/utils/error"
)

type dbRepository struct {
	mysqlClient *sql.DB
}

type DbRepository interface{
	CreateClient(user domain.User) _errors.RestError
}

const (
	queryRegisterUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
)

func NewDbRepository(mysqlClient *sql.DB) DbRepository {
	return &dbRepository{mysqlClient}
}

func (db *dbRepository) CreateClient(user domain.User) _errors.RestError {
	stmt, err := db.mysqlClient.Prepare(queryRegisterUser)
	if err != nil {
		log.Error("error when trying to prepare save user statement", err)
		return _errors.NewInternalServerError("error when trying to save user")
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		log.Error("error when trying to save user", err)
		return _errors.NewInternalServerError("error when trying to save user")
	}

	userId, err := result.LastInsertId()
	if err != nil {
		log.Error("error when trying to get last insert id after creating a new user", err)
		return _errors.NewInternalServerError("error when trying to save user")
	}
	user.ID = userId

	return nil
}
