package cache

import (
	"OpenMall/conf"
	"OpenMall/util/string_util"
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

	if string_util.IsNotNil(err) {
		fmt.Println(err)
	}

	Cache = client
}
