package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kigland/FlashPoint/controller/ping"
	"github.com/kigland/FlashPoint/controller/types"
	"github.com/kigland/FlashPoint/controller/worker"
)

func Init(r gin.IRouter) {
	register(r, &ping.Controller{})
	register(r, &worker.Controller{})
}

func register(r gin.IRouter, cs ...types.IController) {
	for _, c := range cs {
		c.Init(r)
	}
}
