package recovery

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware is a gin.HandlerFunc able to recover
// panics that logs the recovered panic and aborts
// the HTTP request returning an Internal Server Error (500).
func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Recovery from panic
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("[Middleware] %s panic recovered:\n%s\n",
					time.Now().Format("2006/01/02 - 15:04:05"), err)

				ctx.Abort()
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		// Process request
		ctx.Next()
	}
}