package main

import (
	"OpenMall/commands"
	"OpenMall/logger"
	"github.com/spf13/pflag"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
)

var version = "development"

var (
	cfgFile = pflag.StringP("config", "c", "./conf/dev.yaml", "config file path.")
)

func main() {
	//pflag.Parse()
	//cfg := conf.Init()
	//dao.InitMysql(&cfg.Mysql)
	//logger.Init()
	//cronflow.CreateServerDefault(cfg.Machinery, cfg.Redis)
	//cronflow.InitWorker()
	//r := router.InitRouter()
	//err := r.Run(":" + cfg.App.Addr)
	//if err != nil {
	//	panic(err)
	//	return
	//}

	app := cli.NewApp()
	app.Name = "mall"
	app.Version = version
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		commands.StartCommand,
	}
	if err := app.Run(os.Args); err != nil {
		logger.ServiceLogger.Error("main", zap.Error(err))
	}
}
