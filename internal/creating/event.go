package creating

import (
	"context"
	"errors"
	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/marcohb99/go-api-example/internal/increasing"
	event "github.com/marcohb99/go-api-example/kit/events"
)

type IncreaseReleasesCounterOnReleaseCreated struct {
	increasingService increasing.ReleaseCounterService
}

func NewIncreaseReleasesCounterOnReleaseCreated(increaserService increasing.ReleaseCounterService) IncreaseReleasesCounterOnReleaseCreated {
	return IncreaseReleasesCounterOnReleaseCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseReleasesCounterOnReleaseCreated) Handle(_ context.Context, evt event.Event) error {
	releaseCreatedEvt, ok := evt.(apiExample.ReleaseCreatedEvent) // type casting
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(releaseCreatedEvt.ID())
}
