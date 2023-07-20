package cache

import (
	"OpenMall/conf"
	"fmt"
	"github.com/go-redis/redis"
)

var Cache *redis.Client

func InitRedis(conf *conf.Config) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisConf.Host,
		Password: "",
		DB:       conf.RedisConf.Db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err)
	}

	Cache = client
}
