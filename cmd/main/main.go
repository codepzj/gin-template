package main

import (
	"flag"
	"log"

	"git.disoms.net/shared-libs/go-lib/logger"
	"github.com/codepzj/gin-template/conf"
)

// 命令行参数
var (
	CfgPath = flag.String("cfg_path", "./conf/config.yaml", "配置文件路径,eg: ./conf/config.yaml")
)

func main() {
	// 解析命令行参数
	flag.Parse()

	// 根据环境读取配置文件
	config := conf.GetConfig(*CfgPath)

	err := logger.Init(logger.Config{
		Level:  config.Log.Level,
		Format: config.Log.Format,
	})
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	server, err := wireApp(config)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	server.Start()
}
