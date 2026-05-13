package conf

import (
	"log"

	"github.com/codepzj/gin-template/pkg/logger"
	"github.com/spf13/viper"
)

var cfg Config

type Config struct {
	Port  int    `mapstructure:"port"`
	Mysql *Mysql `mapstructure:"mysql"`
	Log   *Log   `mapstructure:"log"`
}

// MySQL配置
type Mysql struct {
	Dsn              string `mapstructure:"dsn"`                // 数据库连接字符串
	SlowSqlThreshold int    `mapstructure:"slow_sql_threshold"` // 慢查询阈值，单位毫秒
	LogLevel         string `mapstructure:"log_level"`          // 日志级别 silent info warn error
}

// Log日志配置
type Log struct {
	Format logger.LogFormat `mapstructure:"format"` // 环境 console json
	Level  logger.LogLevel  `mapstructure:"level"`  // 日志级别
	Output struct {
		EnableFile bool   `mapstructure:"enable_file"`
		FilePath   string `mapstructure:"file_path"`
		MaxAge     int    `mapstructure:"max_age"`
	} `mapstructure:"output"`
}

// GetConfig loads and returns the application configuration
// cfgPath: configuration file path
func GetConfig(cfgPath string) *Config {
	v := viper.New()

	// set config file
	v.SetConfigFile(cfgPath)

	// read config
	err := v.ReadInConfig()
	if err != nil {
		log.Println("config file read error", "path", cfgPath, "error", err.Error())
		panic(err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Println("unmarshal config error:", err.Error())
		panic(err)
	}
	log.Println("config loaded successfully")

	cfg = config
	return &cfg
}

func GetGlobalConfig() *Config {
	return &cfg
}
