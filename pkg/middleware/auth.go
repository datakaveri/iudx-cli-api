package middleware

import (
	"fmt"
	"iudx_domain_specific_apis/pkg/configs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyToken := c.Request.Header.Get("x-api-key")

		if apiKeyToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "API key not found"})
			return
		}

		token, _, err := new(jwt.Parser).ParseUnverified(apiKeyToken, jwt.MapClaims{})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Error with JWT Token"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			iid := fmt.Sprint(claims["iid"])
			if configs.GetAPIKey() == iid {
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT Token"})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
	}
}
