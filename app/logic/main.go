package main

import (
	_ "mahjong/app/logic/msg"

	"github.com/zerak/ego/config"
	"github.com/zerak/ego/log"
	"github.com/zerak/ego/service"
)

func main() {
	log.Trace("mahjong logic server trace")
	log.Info("mahjong logic server info")
	log.Debug("mahjong logic server debug")
	log.Warn("mahjong logic server warn")
	log.Error("mahjong logic server error")

	//log.Fatal("mahjong logic server")
	log.Info("mahjong logic server")
	log.Info("config info:%v", config.Opt)
	service.Run(service.NewTcp(config.Server{}))
	log.Fatal("mahjong logic server fatal")
}
