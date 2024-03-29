package release

import (
	"errors"
	"github.com/marcohb99/go-api-example/kit/command"
	"net/http"

	"github.com/gin-gonic/gin"
	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/marcohb99/go-api-example/internal/creating"
)

type createReleaseRequest struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Released    string `json:"released" binding:"required"`
	ResourceUrl string `json:"resource_url" binding:"required"`
	Uri         string `json:"uri" binding:"required"`
	Year        string `json:"year" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request processing
		var req createReleaseRequest

		// json validation
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		_, err := commandBus.Dispatch(ctx, creating.NewReleaseCommand(
			req.ID,
			req.Title,
			req.Released,
			req.ResourceUrl,
			req.Uri,
			req.Year,
		))

		// Return 400 Bad Request if any validation error occurs.
		if err != nil {
			switch {
			case errors.Is(err, apiExample.ErrInvalidReleaseID),
				errors.Is(err, apiExample.ErrEmptyReleaseTitle), errors.Is(err, apiExample.ErrEmptyReleaseResourceUrl),
				errors.Is(err, apiExample.ErrEmptyReleaseURI), errors.Is(err, apiExample.ErrEmptyReleaseYear):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
