package release

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcohb99/go-api-example/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetAll(t *testing.T)  {
	releaseRepository := new(storagemocks.ReleaseRepository)

	expectedCollection := SampleReleaseCollection(5)
	releaseRepository.On("All", mock.Anything).Return(expectedCollection)

	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/releases", GetAllHandler(releaseRepository))

	// t.Run("when throws an error", func(t *testing.T) {
	// 	// GIVEN
	// 	// we make a GET request to /releases
	// 	req, err := http.NewRequest(http.MethodGet, "/releases", bytes.NewBuffer([]byte{}))
	// 	require.NoError(t, err)

	// 	// WHEN: send request
	// 	rec := httptest.NewRecorder()
	// 	r.ServeHTTP(rec, req)

	// 	// THEN
	// 	res := rec.Result()
	// 	defer res.Body.Close()

	// 	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	// })
	
	t.Run("when get returns a collection", func(t *testing.T) {
		// GIVEN
		// we make a GET request to /releases
		req, err := http.NewRequest(http.MethodGet, "/releases", bytes.NewBuffer([]byte{}))
		require.NoError(t, err)

		// WHEN: send request
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// THEN
		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)

		var serializedBody *createReleaseRequest
		err = json.Unmarshal(body, serializedBody)
		require.NoError(t, err)

		assert.NotEmpty(t, serializedBody.ID)
		assert.NotEmpty(t, serializedBody.Released)
		assert.NotEmpty(t, serializedBody.ResourceUrl)
		assert.NotEmpty(t, serializedBody.Title)
		assert.NotEmpty(t, serializedBody.Year)
		assert.NotEmpty(t, serializedBody.Uri)
	})
}