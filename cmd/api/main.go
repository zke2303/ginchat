package main

import (
	"github.com/nanfeng/ginchat/internal/app"
	"github.com/nanfeng/ginchat/internal/config"
	"github.com/nanfeng/ginchat/internal/server"
)

// @title           GinChat API
// @version         1.0
// @description     这是一个使用 Gin 框架开发的聊天应用 API 服务
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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
