package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koding/multiconfig"
)

type Config struct {
	ListenAddress string `default:":55555"`
	HomeUrl       string `default:"/"`
	StaticFS      []string
	PingInterval  int    `default:5`
	RemoteAddress string `default:"/"`
	Accounts      gin.Accounts
}

var (
	config = new(Config)
)

//读取配置文件
func readConfig() {
	m := multiconfig.NewWithPath("config.toml") // supports TOML and JSON
	// Populated the serverConf struct
	m.MustLoad(config) // Check for error
}
