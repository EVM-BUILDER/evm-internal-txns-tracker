package main

import (
	"github.com/internal-tx/api"
	"github.com/internal-tx/cache"
	"github.com/internal-tx/config"
)

const (
	defaultConfigPath = "./conf.json"
)

func main() {
	config.Init(defaultConfigPath)
	cache.NewMCache()

	api.StartServer()
}
