//go:build wireinject
// +build wireinject

package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/wire"
	"github.com/nanfeng/ginchat/internal/config"
	v1 "github.com/nanfeng/ginchat/internal/handler/v1"
	"github.com/nanfeng/ginchat/internal/repository"
	"github.com/nanfeng/ginchat/internal/server"
	"github.com/nanfeng/ginchat/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type App struct {
	DB *gorm.DB
	// 可以后续加其他字段，例如：
	// Config    *config.AppConfig
	// Logger    *zap.Logger
	HTTPServer *http.Server
}

func BuildApp() (*App, error) {
	wire.Build(
		config.Init,
		NewDB,
		server.NewHTTPServer,
		repository.NewUserRepository,
		service.NewUserService,
		v1.NewUserHandler,
		wire.Struct(new(App), "*"),
	)

	return nil, nil
}

func NewDB(cfg *config.AppConfig) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(cfg.Database.Dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
