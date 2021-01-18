package db

import (
	"database/sql"

	"memesdotcom-users/domain"
	"memesdotcom-users/utils/constants"

	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	log "github.com/sirupsen/logrus"
)

type dbRepository struct {
	mysqlClient *sql.DB
}

type DbRepository interface {
	CreateClient(user domain.User) _errors.RestError
	FindByCredentials(userCredentials *domain.UserCredentials) (*domain.User, _errors.RestError)
}

const (
	queryRegisterUser              = "INSERT INTO users(first_name, last_name, email, username, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?, ?);"
	queryFindByEmailAndPassword    = "SELECT id, first_name, last_name, email, username, date_created FROM users WHERE email=? AND password=? AND status=?;"
	queryFindByUsernameAndPassword = "SELECT id, first_name, last_name, email, username, date_created FROM users WHERE username=? AND password=? AND status=?;"
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

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Username, user.DateCreated, user.Status, user.Password)
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

func (db *dbRepository) FindByCredentials(userCredentials *domain.UserCredentials) (*domain.User, _errors.RestError) {

	var stmt *sql.Stmt
	var err error
	var entry string

	if userCredentials.Email != "" {
		stmt, err = db.mysqlClient.Prepare(queryFindByEmailAndPassword)
		entry = userCredentials.Email
	} else {
		stmt, err = db.mysqlClient.Prepare(queryFindByUsernameAndPassword)
		entry = userCredentials.Username
	}

	if err != nil {
		log.Error("error when trying to prepare search user statement", err)
		return nil, _errors.NewInternalServerError("error when trying to search user")
	}
	defer stmt.Close()

	var user domain.User

	result := stmt.QueryRow(entry, userCredentials.Password, constants.StatusActive)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.DateCreated); err != nil {
		log.Error("error when trying to search user", err)
		return nil, _errors.NewInternalServerError("error when trying to search user")
	}

	return &user, nil
}
