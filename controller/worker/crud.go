package worker

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kigland/FlashPoint/apimod"
	"github.com/kigland/FlashPoint/lib/flashmap"
	"github.com/kigland/FlashPoint/shared"
)

func setCache(c *gin.Context) {
	var req apimod.SetCacheReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("[SET]", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := flashmap.Type(req.Type)
	if t == "" {
		t = flashmap.TypeText
	}

	shared.Cache.Set(req.Key, req.Value, time.Duration(req.TTL)*time.Second, t)
	log.Println("[SET]", req.Key, req.TTL, t)

	c.JSON(http.StatusOK, apimod.SetCacheResp{Key: req.Key})
}
