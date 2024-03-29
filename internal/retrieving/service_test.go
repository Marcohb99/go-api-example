package retrieving

import (
	"context"
	"errors"
	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/marcohb99/go-api-example/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_ReleaseService_GetAllReleases_RepositoryError(t *testing.T) {
	// given a limit
	limit := 10

	// and a release repository that returns an error
	releaseRepository := new(storagemocks.ReleaseRepository)
	releaseRepository.On("GetAll", mock.Anything, limit).Return([]apiExample.Release{}, errors.New("unexpected error"))

	sut := NewReleaseService(releaseRepository)

	// when getting all releases
	releases, err := sut.AllReleases(context.Background(), limit)

	// then assert that the repository was called
	releaseRepository.AssertExpectations(t)

	// and assert that the repository returned an error
	assert.Error(t, err)
	assert.Emptyf(t, releases, "expected no releases, got %v", releases)
}

func Test_ReleaseService_GetAllReleases_Success(t *testing.T) {
	// given a limit
	limit := 10
	r1, _ := apiExample.NewRelease(
		"8a1c5cdc-ba57-445a-994d-aa412d23723f",
		"Ultra Mono",
		"2020-01-01",
		"https://api.discogs.com/releases/1809205",
		"https://www.discogs.com/master/1809205-Idles-Ultra-Mono",
		"2020",
	)
	r2, _ := apiExample.NewRelease(
		"8a1c5cdc-ba57-445a-994e-aa412d23724a",
		"Origin of Symmetry",
		"2001-01-01",
		"https://api.discogs.com/releases/11019",
		"https://www.discogs.com/master/11019-Muse-Origin-of-Symmetry",
		"2001",
	)

	releases := []apiExample.Release{r1, r2}

	// and a release repository that returns an error
	releaseRepository := new(storagemocks.ReleaseRepository)
	releaseRepository.On("GetAll", mock.Anything, limit).Return(releases, nil)

	sut := NewReleaseService(releaseRepository)

	// when getting all releases
	result, err := sut.AllReleases(context.Background(), limit)

	// then assert that the repository was called
	releaseRepository.AssertExpectations(t)

	// and assert that the repository returned an error
	assert.NoError(t, err)
	assert.Equal(t, releases, result, "expected %v, got %v", releases, result)
}
