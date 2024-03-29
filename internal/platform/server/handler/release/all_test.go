package release

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/marcohb99/go-api-example/kit/command/commandmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetAll(t *testing.T) {
	// setup
	// setup
	commandBus := new(commandmocks.Bus)
	r1, _ := apiExample.NewRelease(
		"bad92bf5-9176-47bd-bcc6-8c38a5394d6e",
		"title",
		"2020-01-01",
		"http://resource.com",
		"http://uri.com",
		"2020",
	)
	r2, _ := apiExample.NewRelease(
		"bad92bf5-9176-47bd-bcc6-8c38a5394d6f",
		"title",
		"2020-01-01",
		"http://resource.com",
		"http://uri.com",
		"2020",
	)
	releases := []apiExample.Release{r1, r2}

	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("retrieving.ReleaseCommand"),
	).Return(releases, nil)

	// gin setup
	gin.SetMode(gin.TestMode)
	r := gin.New() // created with no parameters

	// declare routes
	r.GET("/releases", GetAllHandler(commandBus))

	// TESTS
	t.Run("given a valid request it returns 200", func(t *testing.T) {
		// GIVEN
		rr1 := getReleaseResponse{
			ID:          "bad92bf5-9176-47bd-bcc6-8c38a5394d6e",
			Title:       "title",
			Released:    "2020-01-01",
			ResourceUrl: "http://resource.com",
			URI:         "http://uri.com",
			Year:        "2020",
		}
		rr2 := getReleaseResponse{
			ID:          "bad92bf5-9176-47bd-bcc6-8c38a5394d6f",
			Title:       "title",
			Released:    "2020-01-01",
			ResourceUrl: "http://resource.com",
			URI:         "http://uri.com",
			Year:        "2020",
		}
		expectedReleases := []getReleaseResponse{rr1, rr2}

		req, err := http.NewRequest(http.MethodGet, "/releases", bytes.NewBuffer(nil))
		require.NoError(t, err)

		// WHEN: send request
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// THEN: get response and assert data
		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

		data := rec.Body.Bytes()

		var releasesResponse []getReleaseResponse
		err = json.Unmarshal(data, &releasesResponse)
		require.NoError(t, err)

		assert.NotNil(t, releasesResponse)
		assert.Equal(t, expectedReleases, releasesResponse)
	})
}
