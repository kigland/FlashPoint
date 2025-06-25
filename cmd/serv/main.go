package main

import (
	"log"
	"time"

	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"github.com/kigland/FlashPoint/controller"
	"github.com/kigland/FlashPoint/shared"
)

func initCfg() {
	bs, err := iox.ReadAllByte("config.json")
	panicx.NotNilErr(err)
	panicx.NotNilErr(shared.LoadConfig(bs))
}

var gcInterval = time.Minute * 10
var gcTicker = time.NewTicker(gcInterval)

func main() {
	initCfg()
	shared.Init()

	controller.Init(shared.Engine)

	go func() {
		log.Println("[GC] start lazy GC daemon with interval " + gcInterval.String())
		for {
			<-gcTicker.C
			shared.Cache.GC()
		}
	}()

	shared.RunGin()
}
