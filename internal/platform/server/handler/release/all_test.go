package release

import (
	"bytes"
	"github.com/gin-gonic/gin"
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
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("retrieving.ReleaseCommand"),
	).Return(nil)

	// gin setup
	gin.SetMode(gin.TestMode)
	r := gin.New() // created with no parameters

	// declare routes
	r.GET("/releases", GetAllHandler(commandBus))

	// TESTS
	t.Run("given a valid request it returns 200", func(t *testing.T) {
		// GIVEN

		req, err := http.NewRequest(http.MethodGet, "/releases", bytes.NewBuffer(nil))
		require.NoError(t, err)

		// WHEN: send request
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// THEN: get response and assert
		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
