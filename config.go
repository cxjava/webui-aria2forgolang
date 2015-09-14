package main

import "github.com/koding/multiconfig"

type Config struct {
	ListenAddress string `default:":55555"`
	HomeUrl       string `default:"/"`
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
