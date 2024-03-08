package creating

import (
	"context"
	"errors"
	"testing"

	"github.com/marcohb99/go-api-example/internal/platform/storage/storagemocks"
	"github.com/marcohb99/go-api-example/kit/events/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ReleaseService_CreateRelease_RepositoryError(t *testing.T) {
	// given a release
	id := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	title := "Ultra Mono"
	released := "2020-01-01"
	resourceUrl := "https://api.discogs.com/releases/1809205"
	uri := "https://www.discogs.com/master/1809205-Idles-Ultra-Mono"
	year := "2020"

	// and a release repository that returns no error
	releaseRepository := new(storagemocks.ReleaseRepository)
	releaseRepository.On("Save", mock.Anything, mock.AnythingOfType("apiExample.Release")).Return(errors.New("unexpected error"))

	eventBusMock := new(eventmocks.Bus)

	createReleaseSrv := NewReleaseService(releaseRepository, eventBusMock)

	// when creating a release
	err := createReleaseSrv.CreateRelease(context.Background(), id, title, released, resourceUrl, uri, year)

	// then assert that the repository was called and that the bus was not called
	releaseRepository.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
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

	// and a release repository that returns an error
	releaseRepository := new(storagemocks.ReleaseRepository)
	releaseRepository.On("Save", mock.Anything, mock.AnythingOfType("apiExample.Release")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	createReleaseSrv := NewReleaseService(releaseRepository, eventBusMock)

	// when creating a release
	err := createReleaseSrv.CreateRelease(context.Background(), id, title, released, resourceUrl, uri, year)

	// then assert that the repository was called and that the bus was called
	releaseRepository.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	// and assert that the repository returned an error
	assert.NoError(t, err)
}

func Test_ReleaseService_CreateRelease_EventBusError(t *testing.T) {
	// given a release
	id := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	title := "Ultra Mono"
	released := "2020-01-01"
	resourceUrl := "https://api.discogs.com/releases/1809205"
	uri := "https://www.discogs.com/master/1809205-Idles-Ultra-Mono"
	year := "2020"

	// and a release repository that returns an error
	releaseRepository := new(storagemocks.ReleaseRepository)
	releaseRepository.On("Save", mock.Anything, mock.AnythingOfType("apiExample.Release")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something unexpected happened"))

	createReleaseSrv := NewReleaseService(releaseRepository, eventBusMock)

	// when creating a release
	err := createReleaseSrv.CreateRelease(context.Background(), id, title, released, resourceUrl, uri, year)

	// then assert that the repository was called and that the bus was called
	releaseRepository.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	// and assert that the repository returned an error
	assert.Error(t, err)
}
