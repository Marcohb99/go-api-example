package release

import (
	"github.com/gin-gonic/gin"
	"github.com/marcohb99/go-api-example/internal/retrieving"
	"github.com/marcohb99/go-api-example/kit/command"
	"net/http"
)

type getReleasesRequest struct {
	Limit int `json:"limit"`
}

const defaultLimit = 10

func GetAllHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request processing
		var req getReleasesRequest

		var limit int
		limitParam := ctx.Param("limit")

		if limitParam == "" {
			limit = defaultLimit
		} else {
			// json validation
			if err := ctx.BindJSON(&req); err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
		}

		err := commandBus.Dispatch(ctx, retrieving.NewReleaseCommand(limit))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusOK)
	}
}
