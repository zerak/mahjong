package main

import (
	"github.com/zerak/ego/config"
	"github.com/zerak/ego/log"
	"github.com/zerak/ego/service"
)

func main() {
	log.Info("mahjong gate server")
	service.Run(service.NewTcp(config.Server{}))
}
