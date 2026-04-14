package service

import "github.com/codepzj/gin-template/internal/repository"

type Service struct {
	data *repository.Data
}

func NewService(data *repository.Data) *Service {
	return &Service{data: data}
}
