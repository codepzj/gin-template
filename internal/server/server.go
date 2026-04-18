package server

import (
	"fmt"
	"log"

	"github.com/codepzj/gin-template/conf"
	"github.com/codepzj/gin-template/internal/handler"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	cfg       *conf.Config
	ginEngine *GinEngine
}

type GinEngine struct {
	engine *gin.Engine
}

func NewGin(h *handler.Handler) *GinEngine {
	engine := gin.Default()

	h.Register(engine)

	return &GinEngine{
		engine: engine,
	}
}

func (g *GinEngine) Run(port int) error {
	return g.engine.Run(fmt.Sprintf(":%d", port))
}

func NewHttpServer(ginEngine *GinEngine, cfg *conf.Config) *HttpServer {
	return &HttpServer{
		cfg:       cfg,
		ginEngine: ginEngine,
	}
}

func (s *HttpServer) Start() {
	// Start server
	if err := s.ginEngine.Run(s.cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
