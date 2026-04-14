package repository

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
	// 注入数据库、缓存、配置等
}

func NewData() *Data {
	return &Data{}
}
