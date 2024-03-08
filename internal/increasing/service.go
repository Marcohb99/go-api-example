package increasing

type ReleaseCounterService struct{}

func NewReleaseCounterService() ReleaseCounterService {
	return ReleaseCounterService{}
}

func (s ReleaseCounterService) Increase(id string) error {
	return nil
}
