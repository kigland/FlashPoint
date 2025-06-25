package worker

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/kigland/FlashPoint/shared"
)

func getHeadersUntilFound(c *gin.Context, keys ...string) string {
	for _, key := range keys {
		if value := c.GetHeader(key); value != "" {
			return value
		}
	}
	return ""
}

func MidACL(c *gin.Context) {
	apiKey := getHeadersUntilFound(c, "X-API-Key", "X-API-Token", "Authorization")
	if apiKey == "" || !slices.Contains(shared.GetConfig().APIKeys, apiKey) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
}
