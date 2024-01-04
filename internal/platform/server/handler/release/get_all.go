package release

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiExample "github.com/marcohb99/go-api-example/internal"
)

func GetAllHandler(releaseRepository apiExample.ReleaseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := releaseRepository.All(ctx)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, result)
	}
}