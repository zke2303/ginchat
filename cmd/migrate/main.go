package main

import (
	"github.com/nanfeng/ginchat/internal/app"
	"github.com/nanfeng/ginchat/internal/config"
	"github.com/nanfeng/ginchat/internal/model"
)

func main() {
	// 1.加载配置文件
	config.Init()
	// 2.依赖注入
	a, _ := app.BuildApp()

	a.DB.AutoMigrate(&model.User{})
}
