package main

import (
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

func main() {
	initCfg()
	shared.Init()

	controller.Init(shared.Engine)

	go func() {
		for {
			time.Sleep(time.Minute * 10)
			shared.Cache.GC()
		}
	}()

	shared.RunGin()
}
