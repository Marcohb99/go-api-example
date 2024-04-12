package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Middleware is a gin.HandlerFunc that logs some information
// of the incoming request and the consequent response.
func Middleware(apiKeysStr string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKeys := strings.Split(apiKeysStr, ",")
		clientApiKey := ctx.GetHeader("X-API-KEY")
		fmt.Println("API Keys: ", apiKeys)
		fmt.Println("Client API Key: ", clientApiKey)
		if len(clientApiKey) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		found := false

		for _, apiKey := range apiKeys {
			if apiKey == clientApiKey {
				found = true
				break
			}
		}
		if found {
			ctx.Next()
			return
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
