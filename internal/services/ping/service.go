package ping

import (
	"context"

	pingmodel "coupon-system/internal/entity/ping"
	pingrepo "coupon-system/internal/repositories/ping"
)

//go:generate mockgen -source=service.go -package=mock -destination=mock/service_mock.go
type PingServiceProvider interface {
	Ping(ctx context.Context) (pingmodel.PingPong, error)
}

type pingService struct {
	pingRepo pingrepo.PingRepositoryProvider
}

func NewPingService(
	pingRepo pingrepo.PingRepositoryProvider,
) PingServiceProvider {
	return &pingService{
		pingRepo: pingRepo,
	}
}

func (s *pingService) Ping(ctx context.Context) (pingmodel.PingPong, error) {
	err := s.pingRepo.Ping(ctx)
	if err != nil {
		return pingmodel.PingPong{
			Message: pingmodel.ErrorMessage,
		}, err
	}

	return pingmodel.PingPong{
		Message: pingmodel.SuccessMessage,
	}, nil
}