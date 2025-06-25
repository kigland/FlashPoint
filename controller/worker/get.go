package worker

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kigland/FlashPoint/lib/flashmap"
	"github.com/kigland/FlashPoint/shared"
)

func getCache(c *gin.Context) {
	isRaw := strings.ToLower(c.Query("raw")) == "true"

	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key is required"})
		return
	}

	value, ok := shared.Cache.Get(key)
	if !ok || value.Value == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	log.Println("[GET]", "key="+key, "isRaw="+strconv.FormatBool(isRaw))

	if isRaw {
		c.JSON(http.StatusOK, value)
		return
	}

	if value.Mime != "" {
		c.Header("Content-Type", value.Mime)
	}

	switch value.Type {
	case flashmap.TypeBinary:
		switch vT := value.Value.(type) {
		case []byte:
			c.Data(http.StatusOK, value.Mime, vT)
		case string:
			c.Data(http.StatusOK, value.Mime, []byte(vT))
		default:
			c.JSON(http.StatusOK, value.Value)
		}
	case flashmap.TypeText:
		c.String(http.StatusOK, value.Value.(string))
	case flashmap.TypeJSON:
		c.JSON(http.StatusOK, value.Value)
	default:
		c.JSON(http.StatusOK, value.Value)
	}
}
