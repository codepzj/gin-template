package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codepzj/gin-template/conf"
	"github.com/codepzj/gin-template/internal/handler"
	"github.com/codepzj/gin-template/internal/service"
	"github.com/codepzj/gin-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

	// 先初始化日志
	logger.NewLogger(&logger.Option{
		Format: config.Log.Format,
		Level:  config.Log.Level,
		Output: &logger.OutputConfig{
			EnableFile: config.Log.Output.EnableFile,
			FilePath:   config.Log.Output.FilePath,
			MaxAge:     config.Log.Output.MaxAge,
		},
	})
	defer logger.Sync()

	// 初始化依赖
	healthSvc := service.NewHealthService()
	healthHdl := handler.NewHealthHandler(healthSvc)

	// 注册路由
	engine := gin.Default()
	healthHdl.Register(engine)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.GetLogger().Fatal("listen error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown", zap.Error(err))
	}
	logger.Info("server stopped")
}
