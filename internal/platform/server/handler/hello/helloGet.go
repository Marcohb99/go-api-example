package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		ctx.String(http.StatusOK, "hello")
	}
}