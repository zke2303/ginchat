package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/wire"
	"github.com/nanfeng/ginchat/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func BuildApp() (*App, error) {
	wire.Build(
		config.Init,
		NewDB,
	)

	return nil, nil
}

type App struct {
	HTTPServer *http.Server
}

func NewDB() *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(config.Cfg.Database.Dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("数据库连接失败, cause: " + err.Error())
	}

	return db
}
