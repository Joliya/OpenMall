package cache

import (
	"OpenMall/conf"
	"github.com/go-redis/redis"
)

func InitRedis(conf *conf.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisConf.Host,
		Password: conf.RedisConf.Password,
		DB:       conf.RedisConf.Db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	return client
}
