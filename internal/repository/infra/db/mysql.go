package db

import (
	"log"
	"os"
	"time"

	"github.com/codepzj/gin-template/conf"
	"github.com/codepzj/go-lib/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDB(cfg *conf.Mysql) *gorm.DB {
	newLogger := glogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		glogger.Config{
			SlowThreshold: time.Duration(cfg.SlowSqlThreshold) * time.Millisecond, // 慢查询阈值，单位毫秒
			LogLevel:      parseMysqlLogLevel(cfg),                                // 日志级别
		},
	)
	// 初始化gorm
	db, err := gorm.Open(mysql.New(
		mysql.Config{DSN: cfg.Dsn}),
		&gorm.Config{Logger: newLogger, NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}},
	)
	if err != nil {
		logger.Error("Connect MySQL failed", zap.Error(err))
		panic(err)
	}

	logger.Info("MySQL connect success...")
	return db
}

func parseMysqlLogLevel(cfg *conf.Mysql) glogger.LogLevel {
	switch cfg.LogLevel {
	case "silent":
		return glogger.Silent
	case "info":
		return glogger.Info
	case "warn":
		return glogger.Warn
	case "error":
		return glogger.Error
	}
	return glogger.Info
}
