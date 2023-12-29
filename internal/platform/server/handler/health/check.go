package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckHandler returns an HTTP handler to perform health checks.
// 2nd param ignore because hc endpoints do not need request data
func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "everything is ok!")
	}
}