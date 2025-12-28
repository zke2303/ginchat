package utils

import "github.com/nanfeng/ginchat/internal/config"

var Secret = []byte(config.Cfg.Secret)

type Claims struct {
	Id       string
	username string
}

func GenerateToken(id string, username string) (string, error) {

}

func ParseToken(token string) (Claims, error) {

}
