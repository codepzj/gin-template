package handler

import (
	"net/http"

	"github.com/codepzj/gin-template/internal/service"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	svc *service.HealthService
}

func NewHealthHandler(svc *service.HealthService) *HealthHandler {
	return &HealthHandler{svc: svc}
}

func (h *HealthHandler) Register(engine *gin.Engine) {
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
