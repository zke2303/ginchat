package main

import (
	"github.com/nanfeng/ginchat/internal/app"
	"github.com/nanfeng/ginchat/internal/config"
	"github.com/nanfeng/ginchat/internal/server"
)

// gin-swagger middleware
// swagger embed files
func main() {
	// 1.加载配置文件
	config.Init()
	// 2.依赖注入
	a, err := app.BuildApp()
	if err != nil {
		panic(err)
	}

	server.RunHttpServer(a.HTTPServer)
}
