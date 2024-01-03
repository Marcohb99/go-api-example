package release

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiExample "github.com/marcohb99/go-api-example/internal"
)

const (
	dbUser = "mhb"
	dbPass = "root"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "go_api_example"
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
func CreateHandler(releaseRepository apiExample.ReleaseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request processing
		var req createReleaseRequest

		// json validation
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// instantiate object
		release := apiExample.NewRelease(req.ID, req.Title, req.Released, req.ResourceUrl, req.Uri, req.Year)

		// save object
		if err := releaseRepository.Save(ctx, release); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}