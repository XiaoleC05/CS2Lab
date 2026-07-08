package handler

import (
	"fmt"
	"net/http"

	"github.com/XiaoleC05/CS2Lab/internal/config"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware handles OXELIA_GATEWAY_MODE authentication
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !cfg.OxeliaGatewayMode {
			// In development mode, use a test user ID
			c.Set("userID", int64(1))
			c.Next()
			return
		}

		// In gateway mode, extract user ID from Oxelia51 gateway headers
		userIDStr := c.GetHeader("X-User-ID")
		if userIDStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing X-User-ID header"})
			c.Abort()
			return
		}

		var userID int64
		_, err := fmt.Sscanf(userIDStr, "%d", &userID)
		if err != nil || userID <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid X-User-ID header"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
