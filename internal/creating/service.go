package creating

import (
	"context"

	apiExample "github.com/marcohb99/go-api-example/internal"
	event "github.com/marcohb99/go-api-example/kit/events"
)

// ReleaseService is the default ReleaseService interface
// implementation returned by creating.NewReleaseService.
type ReleaseService struct {
	releaseRepository apiExample.ReleaseRepository
	eventBus         event.Bus
}

// NewReleaseService returns the default Service interface implementation.
func NewReleaseService(releaseRepository apiExample.ReleaseRepository, eventBus event.Bus) ReleaseService {
	return ReleaseService{
		releaseRepository: releaseRepository,
		eventBus: eventBus,
	}
}

func (s ReleaseService) CreateRelease(ctx context.Context, id, title, released, resourceUrl, uri, year string) error {
	release, err := apiExample.NewRelease(id, title, released, resourceUrl, uri, year)
	if err != nil {
		return err
	}
	
	if err := s.releaseRepository.Save(ctx, release); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, release.PullEvents())
}