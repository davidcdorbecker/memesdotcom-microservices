package app

import (
	"context"
	"github.com/go-resty/resty/v2"
	"memesdotcom-auth/handlers"
	redis2 "memesdotcom-auth/infrastructure/repository/redis"
	"memesdotcom-auth/infrastructure/repository/users_api"
	"memesdotcom-auth/infrastructure/router"
	"memesdotcom-auth/services"
	"memesdotcom-auth/utils/constants"

	"github.com/go-redis/redis/v8"
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	handler := handlers.NewAuthHandler(services.NewAuthService(users_api.NewUsersAPI(resty.New()), redis2.NewRedisRepo(redisClient, context.Background())))

	//getting API
	app := router.CreateRestRouter(handler)
	if err := app.Listen(viper.GetString(constants.ApplicationPort)); err != nil {
		panic(err)
	}
}
