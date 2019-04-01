package library

import (
	"fmt"

	"github.com/fuyi-huang/backend/config"
	extend "github.com/fuyi-huang/backend/ext"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func ConnectRedis() uint {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis["host"] + ":" + config.Redis["port"],
		Password: config.Redis["passwd"],
		DB:       0,
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		return extend.REDIS_OP_ERR
	}
	fmt.Println(pong)
	return extend.SUCCESS
}
