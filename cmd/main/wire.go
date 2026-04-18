//go:build wireinject
// +build wireinject

package main

import (
	"github.com/codepzj/gin-template/conf"
	"github.com/codepzj/gin-template/internal/handler"
	"github.com/codepzj/gin-template/internal/repository"
	"github.com/codepzj/gin-template/internal/server"
	"github.com/codepzj/gin-template/internal/service"
	"github.com/google/wire"
)

// wireApp 初始化应用
func wireApp(cfg *conf.Config) (*server.HttpServer, error) {
	panic(wire.Build(
		server.ProviderSet,
		handler.ProviderSet,
		service.ProviderSet,
		repository.ProviderSet,
	))
}
