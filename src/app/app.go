package app

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"memesdotcom-users/handlers"
	"memesdotcom-users/infrastructure/repository/db"
	"memesdotcom-users/infrastructure/router"
	"memesdotcom-users/services"
	"memesdotcom-users/utils/constants"
)

func StartApp() {
	//getting application env variables
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./src")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("config file not found")
	}

	log.Printf("config keys: %s\n", viper.AllKeys())

	//getting mysql config

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		viper.GetString(constants.MySQLUsername),
		viper.GetString(constants.MySQLPassword),
		viper.GetString(constants.MySQLHost),
		viper.GetString(constants.MySQLSchema),
	)

	mysqlClient, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = mysqlClient.Ping(); err != nil {
		panic(err)
	}

	if err = mysql.SetLogger(log.New()); err != nil {
		panic(err)
	}

	log.Infoln("mysql correctly configured")

	//dependency injection
	handler := handlers.NewUsersHandler(services.NewService(db.NewDbRepository(mysqlClient)))

	//getting API
	app := router.CreateRestRouter(handler)
	if err := app.Listen(viper.GetString(constants.ApplicationPort)); err != nil {
		panic(err)
	}
}
