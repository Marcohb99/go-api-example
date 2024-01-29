package creating

import (
	"context"
	"errors"
	"testing"

	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/marcohb99/go-api-example/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_ReleaseService_CreateRelease_RepositoryError(t *testing.T) {
	// given a release
	id := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	title := "Ultra Mono"
	released := "2020-01-01"
	resourceUrl := "https://api.discogs.com/releases/1809205"
	uri := "https://www.discogs.com/master/1809205-Idles-Ultra-Mono"
	year := "2020"
	release, err := apiExample.NewRelease(id, title, released, resourceUrl, uri, "2020")
	require.NoError(t, err)

	// and a release repository that returns no error
	releaseRepository := new(storagemocks.ReleaseRepository)
	releaseRepository.On("Save", mock.Anything, release).Return(errors.New("unexpected error"))

	createReleaseSrv := NewReleaseService(releaseRepository)

	// when creating a release
	err = createReleaseSrv.CreateRelease(context.Background(), id, title, released, resourceUrl, uri, year)

	// then assert that the repository was called
	releaseRepository.AssertExpectations(t)
	// and assert that the repository returned no errors
	assert.Error(t, err)
}

func Test_ReleaseService_CreateRelease_Success(t *testing.T) {
	// given a release
	id := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	title := "Ultra Mono"
	released := "2020-01-01"
	resourceUrl := "https://api.discogs.com/releases/1809205"
	uri := "https://www.discogs.com/master/1809205-Idles-Ultra-Mono"
	year := "2020"
	release, err := apiExample.NewRelease(id, title, released, resourceUrl, uri, "2020")
	require.NoError(t, err)

	// and a release repository that returns an error
	releaseRepository := new(storagemocks.ReleaseRepository)
	releaseRepository.On("Save", mock.Anything, release).Return(nil)

	createReleaseSrv := NewReleaseService(releaseRepository)

	// when creating a release
	err = createReleaseSrv.CreateRelease(context.Background(), id, title, released, resourceUrl, uri, year)

	// then assert that the repository was called
	releaseRepository.AssertExpectations(t)
	// and assert that the repository returned an error
	assert.NoError(t, err)
}
