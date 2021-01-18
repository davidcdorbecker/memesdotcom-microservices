package redis

import (
	"context"
	_errors "github.com/davidcdorbecker/memesdotcom-microservices/memesdotcom-utils/error"
	_redis "github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"memesdotcom-auth/utils/constants"
	"time"
)

type redis struct {
	redisClient *_redis.Client
	ctx         context.Context
}

type Redis interface {
	Get(key string) (string, _errors.RestError)
	Set(key, value string, expirationTime time.Duration) _errors.RestError
	Delete(key string) _errors.RestError
}

func NewRedisRepo(redisClient *_redis.Client, ctx context.Context) Redis {
	return &redis{redisClient, ctx}
}

func (r *redis) Get(key string) (string, _errors.RestError) {
	res, err := r.redisClient.Get(r.ctx, key).Result()
	if err != nil {
		log.Error(err.Error())
		return "", _errors.NewInternalServerError(constants.AccessTokenError)
	}
	return res, nil
}

func (r *redis) Set(key, value string, expirationTime time.Duration) _errors.RestError {
	err := r.redisClient.Set(r.ctx, key, value, expirationTime).Err()
	if err != nil {
		log.Error(err.Error())
		return _errors.NewInternalServerError(constants.AccessTokenError)
	}
	return nil
}

func (r *redis) Delete(key string) _errors.RestError {
	err := r.redisClient.Del(r.ctx, key).Err()
	if err != nil {
		log.Error(err.Error())
		return _errors.NewInternalServerError(constants.AccessTokenError)
	}
	return nil
}
