package release

import (
	"github.com/gin-gonic/gin"
	"github.com/marcohb99/go-api-example/internal/retrieving"
	command "github.com/marcohb99/go-api-example/kit"
	"net/http"
)

func GetAllHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		err := commandBus.Dispatch(ctx, retrieving.NewReleaseCommand())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusOK)
	}
}
