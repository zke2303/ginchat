package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg *AppConfig

type AppConfig struct {
	App      App      `mapstructure:"app"`
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Secret   string   `mapstructure:"secret"`
}

type App struct {
	Name string `mapstructure:"name"`
	mode string `mapstructure:"mode"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Database struct {
	Dsn string `mapstructure:"dsn"`
}

func Init() *AppConfig {
	// 1.创建一个viper对象
	v := viper.New()

	// 2.指定 配置 文件路径和后缀名
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	// 3.读取配置文件信息
	if err := v.ReadInConfig(); err != nil {
		panic("读取配置文件失败, err: " + err.Error())
	}

	// 4.解析配置文件
	if err := v.Unmarshal(&Cfg); err != nil {
		panic("配置文件解析失败, err: " + err.Error())
	}

	// 5.开启热加载
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置热更新")
		if err := v.Unmarshal(Cfg); err != nil {
			panic("配置热更新失败, err: " + err.Error())
		}
	})

	return Cfg
}
