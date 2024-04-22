package commands

import (
	"OpenMall/conf"
	"OpenMall/db/dao"
	"OpenMall/logger"
	"OpenMall/routers"
	"github.com/urfave/cli"
)

var StartCommand = cli.Command{
	Name:  "start",
	Usage: "Starts web server",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "port",
			Value: 9091,
			Usage: "指定服务端口",
		},
		&cli.StringFlag{
			Name:  "env",
			Value: "dev",
			Usage: "指定当前环境， 可选值：dev, test, prod， 默认 dev",
		},
	},
	Action: start,
}

func start(ctx *cli.Context) {
	// go run cmd/main.go start conf/dev.yaml port
	cfg := conf.Init(ctx.String("env"))
	dao.InitMysql(&cfg.Mysql)
	logger.Init()
	r := routers.InitRouter()
	port := ctx.String("port")
	if port == "" {
		port = cfg.App.Addr
	}
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
		return
	}
}
