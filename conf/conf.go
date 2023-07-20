package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppConf   *AppConfig
	SqlConn   string
	RedisConf *RedisConfig
	//....
}

type AppConfig struct {
	Name string
	Port string
	Mode string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
}

func InitConfig() *Config {
	viper.SetConfigType("yaml") //设置配置文件格式
	viper.AddConfigPath("conf") //设置配置文件的路径
	viper.SetConfigName("app")  //设置配置文件名
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	//返回配置文件的数据
	return &Config{
		AppConf: &AppConfig{
			Name: viper.GetString("app.name"),
			Port: viper.GetString("app.port"),
			Mode: viper.GetString("app.mode"),
		},
		SqlConn: viper.GetString("mysql.conn"),
		RedisConf: &RedisConfig{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetString("redis.port"),
			Password: viper.GetString("redis.password"),
			Db:       viper.GetInt("redis.db"),
		},
	}
}
