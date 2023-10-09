package middleware

import (
	"iudx_domain_specific_apis/pkg/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyToken := c.Request.Header.Get("x-api-key")

		if apiKeyToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "API key not found"})
			return
		}

		if configs.GetAPIKey() == apiKeyToken {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
	}
}
