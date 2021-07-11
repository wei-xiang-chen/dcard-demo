package client

import (
	"dcard/setting"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	RedisEngine *redis.Client
)

func NewRedisEngine(redisSetting *setting.RedisSetting) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisSetting.Addr,
		Password: redisSetting.Password, // no password set
		DB:       redisSetting.DB,       // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
