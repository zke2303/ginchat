package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanfeng/ginchat/internal/config"
	v1 "github.com/nanfeng/ginchat/internal/handler/v1"
)

func NewHTTPServer(handler *v1.UserHandler) *http.Server {
	r := gin.Default()

	api := r.Group("/api")

	handler.Register(api)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Cfg.Server.Port),
		Handler: r,
	}
}

func RunHttpServer(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("服务器启动失败, cause: " + err.Error())
		}
	}()
}
