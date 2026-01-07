package coupon

import (
	"context"
	cpmodel "coupon-system/internal/entity/coupon"
	cprepo "coupon-system/internal/repositories/coupon"
)

//go:generate mockgen -source=service.go -package=mock -destination=mock/service_mock.go
type CouponServiceProvider interface {
	InsertCoupon(ctx context.Context, name string, amount int) (int64, error)
	ClaimCoupon(ctx context.Context, userID, couponName string) error
	GetCouponDetailsByName(ctx context.Context, name string) (cpmodel.CouponDetail, error) 
}

type couponService struct {
	repo    cprepo.CouponRepositoryProvider
}

func NewCouponService(
	couponRepo cprepo.CouponRepositoryProvider,
) CouponServiceProvider {
	return &couponService{
		repo:    couponRepo,
	}
}

func (s *couponService) InsertCoupon(ctx context.Context, name string, amount int) (int64, error) {
	couponID, err := s.repo.InsertCoupon(ctx, name, amount)
	if err != nil {
		return 0, err
	}

	return couponID, nil
}

func (s *couponService) ClaimCoupon(ctx context.Context, userID, couponName string) error{
	 err := s.repo.ClaimCoupon(ctx, userID, couponName)
	if err != nil {
		return  err
	}

	return nil
}

func (s *couponService)GetCouponDetailsByName(ctx context.Context, name string) (cpmodel.CouponDetail, error) {
	couponDetails, err := s.repo.GetCouponDetailsByName(ctx, name)
	if err != nil {
		return cpmodel.CouponDetail{}, err
	}

	return couponDetails, nil
}