package worker

import (
	"github.com/gin-gonic/gin"
	"github.com/kigland/FlashPoint/controller/types"
)

type Controller struct{}

var _ types.IController = (*Controller)(nil)

func (c *Controller) Init(r gin.IRouter) {
	r.POST("/set", MidACL, setCache)
	r.GET("/:key", getCache)
}
