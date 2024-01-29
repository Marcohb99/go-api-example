package creating

import (
	"context"

	apiExample "github.com/marcohb99/go-api-example/internal"
)

// ReleaseService is the default ReleaseService interface
// implementation returned by creating.NewReleaseService.
type ReleaseService struct {
	releaseRepository apiExample.ReleaseRepository
}

// NewReleaseService returns the default Service interface implementation.
func NewReleaseService(releaseRepository apiExample.ReleaseRepository) ReleaseService {
	return ReleaseService{
		releaseRepository: releaseRepository,
	}
}

func (s ReleaseService) CreateRelease(ctx context.Context, id, title, released, resourceUrl, uri, year string) error {
	release, err := apiExample.NewRelease(id, title, released, resourceUrl, uri, year)
	if err != nil {
		return err
	}
	return s.releaseRepository.Save(ctx, release)	
}