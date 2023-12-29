package release

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/marcohb99/go-api-example/internal/platform/storage/mysql"
)

const (
	dbUser = "mhb"
	dbPass = "root"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "go_api_example"
)

type createRequest struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Released    string `json:"released" binding:"required"`
	ResourceUrl string `json:"resource_url" binding:"required"`
	Uri         string `json:"uri" binding:"required"`
	Year        string `json:"year" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request processing
		var req createRequest

		// json validation
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// instantiate object
		course := apiExample.NewRelease(req.ID, req.Title, req.Released, req.ResourceUrl, req.Uri, req.Year)

		// database connection
		mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		db, err := sql.Open("mysql", mysqlURI)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// persist instantiating the repository
		courseRepository := mysql.NewReleaseRepository(db)

		if err := courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}