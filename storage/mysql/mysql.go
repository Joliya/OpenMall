package mysql

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Config struct {
	DSN             string
	SlowThreshold   time.Duration
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime int
	ShowLog         bool
}

func NewMySQL(c *Config) (db *gorm.DB) {
	sqlDB, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTime))

	db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), gormConfig(c))
	if err != nil {
		panic(err)
	}
	//db.Set("gorm:table_options", "CHARSET=utf8mb4")
	return db
}

func gormConfig(c *Config) *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止外键约束
	}
	if c.ShowLog {
		config.Logger = logger.Default.LogMode(logger.Info)
	} else {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	//
	// 只打印慢查询
	if c.SlowThreshold > 0 {
		config.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second * c.SlowThreshold,
				Colorful:      true,
				LogLevel:      logger.Warn,
			},
		)
	}
	config.NamingStrategy = schema.NamingStrategy{
		SingularTable: true, // 表名禁止自动复数
	}
	return config
}
