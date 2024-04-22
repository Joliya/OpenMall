/**
 * @Author: jinpeng zhang
 * @Date: 2023/8/13 14:30
 * @Description:
 */

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var ServiceLogger *zap.Logger
var RouterLogger *zap.Logger

func Init() {
	serviceWriteSyncer := getServiceLogWriter()
	routerWriteSyncer := getRouterLogWriter()
	encoder := getEncoder()
	serviceCore := zapcore.NewCore(encoder, serviceWriteSyncer, zapcore.DebugLevel)
	routerCore := zapcore.NewCore(encoder, routerWriteSyncer, zapcore.DebugLevel)
	ServiceLogger = zap.New(serviceCore)
	RouterLogger = zap.New(routerCore)
}

func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "time"
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	return zapcore.NewJSONEncoder(cfg)
}

func getServiceLogWriter() zapcore.WriteSyncer {
	file, err := os.OpenFile("./logs/service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

func getRouterLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./logs/router.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return zapcore.AddSync(file)
}
