package release

import (
	"github.com/gin-gonic/gin"
	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/marcohb99/go-api-example/internal/retrieving"
	"github.com/marcohb99/go-api-example/kit/command"
	"net/http"
)

type getReleasesRequest struct {
	Limit int `json:"limit"`
}

type getReleaseResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Released    string `json:"released"`
	ResourceUrl string `json:"resourceUrl"`
	URI         string `json:"uri"`
	Year        string `json:"year"`
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

		result, err := commandBus.Dispatch(ctx, retrieving.NewReleaseCommand(limit))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		releases := result.([]apiExample.Release)
		ctx.JSON(http.StatusOK, formatResponse(releases))
	}
}

func formatResponse(releases []apiExample.Release) []getReleaseResponse {
	var response []getReleaseResponse
	for _, release := range releases {
		response = append(response, getReleaseResponse{
			ID:          release.ID().String(),
			Title:       release.Title().String(),
			Released:    release.Released().String(),
			ResourceUrl: release.ResourceUrl().String(),
			URI:         release.Uri().String(),
			Year:        release.Year().String(),
		})
	}
	return response
}
