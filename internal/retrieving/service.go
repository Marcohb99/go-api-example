package retrieving

import (
	"context"
	apiExample "github.com/marcohb99/go-api-example/internal"
)

type ReleaseService struct {
	releaseRepository apiExample.ReleaseRepository
}

func NewReleaseService(releaseRepository apiExample.ReleaseRepository) ReleaseService {
	return ReleaseService{
		releaseRepository: releaseRepository,
	}
}

func (s ReleaseService) AllReleases(ctx context.Context, limit int) ([]apiExample.Release, error) {
	return s.releaseRepository.GetAll(ctx, limit)
}
