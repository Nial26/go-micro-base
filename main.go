package main

import (
	"go-micro-base/config"
	"go-micro-base/server"
)

const (
	configName = "go-micro-base.toml"
	configType = "toml"
	configPath = "."
)

func main() {
	c := config.Load(configName, configType, configPath)
	srv := server.New(c)
	srv.Start()
}
