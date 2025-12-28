package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanfeng/ginchat/internal/config"
	v1 "github.com/nanfeng/ginchat/internal/handler/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/nanfeng/ginchat/docs" // 导入生成的 docs 包
)

func NewHTTPServer(handler *v1.UserHandler) *http.Server {
	r := gin.Default()

	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/v1/api")

	handler.Register(api)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Cfg.Server.Port),
		Handler: r,
	}
}

func RunHttpServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic("服务器启动失败, cause: " + err.Error())
	}
}
