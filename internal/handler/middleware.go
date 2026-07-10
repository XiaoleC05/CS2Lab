package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/XiaoleC05/CS2Lab/internal/config"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware handles OXELIA_GATEWAY_MODE authentication.
// In gateway mode, requires valid X-User-Id header from the gateway.
// Test authentication is only available when GIN_MODE=debug AND TEST_AUTH_ENABLED=true.
// Non-gateway mode without test auth enabled returns 401.
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if cfg.OxeliaGatewayMode {
			userIDStr := c.GetHeader("X-User-Id")
			if userIDStr == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				c.Abort()
				return
			}

			var userID int64
			_, err := fmt.Sscanf(userIDStr, "%d", &userID)
			if err != nil || userID <= 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				c.Abort()
				return
			}

			c.Set("userID", userID)
			c.Next()
			return
		}

		// Non-gateway mode: test auth only if GIN_MODE=debug AND TEST_AUTH_ENABLED=true
		if os.Getenv("GIN_MODE") == "debug" && os.Getenv("TEST_AUTH_ENABLED") == "true" {
			testUserID := c.GetHeader("X-Test-User-Id")
			if testUserID != "" {
				var uid int64
				_, err := fmt.Sscanf(testUserID, "%d", &uid)
				if err == nil && uid > 0 {
					c.Set("userID", uid)
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
	}
}
