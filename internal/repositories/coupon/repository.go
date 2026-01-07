package coupon

import (
	"context"
	cpmodel "coupon-system/internal/entity/coupon"
	"coupon-system/internal/postgres"
)

//go:generate mockgen -source=repository.go -package=mock -destination=mock/repository_mock.go
type CouponRepositoryProvider interface {
	InsertCoupon(ctx context.Context, name string, amount int) (int64, error)
	ClaimCoupon(ctx context.Context, userID, couponName string) error
	GetCouponDetailsByName(ctx context.Context, name string) (cpmodel.CouponDetail, error) 
}

type couponRepository struct {
	db    dbRepoProvider
}

func NewCouponRepository(
	dbCoupon *postgres.Postgres,
) CouponRepositoryProvider {
	return &couponRepository{
		db: newDBRepo(
			dbCoupon,
		),
	}
}

func (r *couponRepository) InsertCoupon(ctx context.Context, name string, amount int) (int64, error) {
	id, err := r.db.InsertCoupon(ctx, name, amount)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *couponRepository) ClaimCoupon(ctx context.Context, userID, couponName string) error {
	err := r.db.ClaimCoupon(ctx, userID, couponName)
	if err != nil {
		return err
	}

	return nil
}

func (r *couponRepository) GetCouponDetailsByName(ctx context.Context, name string) (cpmodel.CouponDetail, error)  {
	out, err := r.db.GetCouponDetailsByName(ctx, name)
	if err != nil {
		return cpmodel.CouponDetail{}, nil
	}

	return out, nil
}