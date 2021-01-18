package app

import (
	"github.com/go-resty/resty/v2"
	"memesdotcom-auth/handlers"
	"memesdotcom-auth/infrastructure/repository/users_api"
	"memesdotcom-auth/infrastructure/router"
	"memesdotcom-auth/services"
	"memesdotcom-auth/utils/constants"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	handler := handlers.NewAuthHandler(services.NewAuthService(users_api.NewUsersAPI(resty.New())))

	//getting API
	app := router.CreateRestRouter(handler)
	if err := app.Listen(viper.GetString(constants.ApplicationPort)); err != nil {
		panic(err)
	}
}
