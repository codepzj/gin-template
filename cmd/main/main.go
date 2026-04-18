package main

import (
	"flag"
	"log"

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

	server, err := wireApp(config)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	server.Start()
}
