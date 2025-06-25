package worker

import (
	"encoding/base64"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kigland/FlashPoint/apimod"
	"github.com/kigland/FlashPoint/lib/flashmap"
	"github.com/kigland/FlashPoint/shared"
)

func setCache(c *gin.Context) {
	var req apimod.SetCacheReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("[SET] err=" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	key := req.Key
	key = strings.TrimSpace(key)
	if key == "" {
		key = uuid.New().String()
	}

	t := flashmap.Type(req.Type)
	if t == "" {
		t = flashmap.TypeText
	}

	var finalValue any = req.Value
	if t == flashmap.TypeBinary {
		if strValue, ok := req.Value.(string); ok {
			decoded, err := base64.StdEncoding.DecodeString(strValue)
			if err != nil {
				log.Println("[SET] base64 decode error=" + err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 data"})
				return
			}
			finalValue = decoded
		}
	}

	shared.Cache.Set(key, finalValue, time.Duration(req.TTL)*time.Second, t, req.Mime)
	log.Println("[SET] key=" + key + " ttl=" + strconv.Itoa(req.TTL) + "s" + " type=" + string(t))

	c.JSON(http.StatusOK, apimod.SetCacheResp{Key: key})
}
