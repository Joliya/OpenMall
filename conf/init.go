package conf

import (
	"OpenMall/storage/mysql"
	"OpenMall/storage/redis"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var (
	Conf = &Config{}
)

type Config struct {
	App   AppConfig
	Mysql mysql.Config
	Redis redis.Config
}

type AppConfig struct {
	Name string
	Addr string
}

func Init(env string) *Config {
	if env == "" {
		env := os.Getenv("ENV")
		if env == "" {
			env = "dev"
		}
	}
	fmt.Println(env)
	fileName := env + ".yaml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/" + fileName)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&Conf); err != nil {
			panic(err)
		}
	})
	return Conf
}
