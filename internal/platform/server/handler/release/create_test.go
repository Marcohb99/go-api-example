package release

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcohb99/go-api-example/internal/creating"
	"github.com/marcohb99/go-api-example/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T)  {
	// setup
	releaseRepository := new(storagemocks.ReleaseRepository)

	// when receiving the method save with any context and any release object
	// it will return no errors
	releaseRepository.On("Save", mock.Anything, mock.Anything).Return(nil)

	createReleaseSrv := creating.NewReleaseService(releaseRepository)

	// gin setup
	gin.SetMode(gin.TestMode)
	r := gin.New() // created with no parameters
	
	// declare routes
	r.POST("/releases", CreateHandler(createReleaseSrv))

	// TESTS
	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		// GIVEN
		request := createReleaseRequest{
			ID: "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Title: "Ultra Mono",
			Released: "2020-01-01",
			ResourceUrl: "https://api.discogs.com/releases/1809205",
			Uri: "https://www.discogs.com/master/1809205-Idles-Ultra-Mono",
			// year is missing
		}

		b, err := json.Marshal(request)
		// assert that no errors are thrown after function
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/releases", bytes.NewBuffer(b))
		require.NoError(t, err)
		
		// WHEN: send request
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// THEN: get response and assert
		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		// GIVEN
		request := createReleaseRequest{
			ID: "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Title: "Ultra Mono",
			Released: "2020-01-01",
			ResourceUrl: "https://api.discogs.com/releases/1809205",
			Uri: "https://www.discogs.com/master/1809205-Idles-Ultra-Mono",
			Year: "2020",
		}

		b, err := json.Marshal(request)
		// assert that no errors are thrown after function
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/releases", bytes.NewBuffer(b))
		require.NoError(t, err)
		
		// WHEN: send request
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// THEN: get response and assert
		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}