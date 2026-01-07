package coupon

import (
	"context"
	cpmodel "coupon-system/internal/entity/coupon"
	"coupon-system/internal/postgres"
	"database/sql"
	"log"

	"github.com/lib/pq"
)

//go:generate mockgen -source=repository_db.go -package=mock -destination=mock/repository_db_mock.go
type dbRepoProvider interface {
	InsertCoupon(ctx context.Context, name string, amount int) (int64, error)
	ClaimCoupon(ctx context.Context, userID, couponName string) error
	GetCouponDetailsByName(ctx context.Context, name string) (cpmodel.CouponDetail, error) 
}

type dbRepo struct {
	dbCoupon    *postgres.Postgres
}

func newDBRepo(
	dbCoupon *postgres.Postgres,
) dbRepoProvider {
	return &dbRepo{
		dbCoupon: dbCoupon,
	}
}

func (r *dbRepo) InsertCoupon(ctx context.Context, name string, amount int) (int64, error) {
	var id int64

	err := r.dbCoupon.DB.QueryRowContext(
		ctx,
		queryInsertCoupon,
		name,
		amount,
	).Scan(&id)

	if err != nil {
		log.Printf("error inserting coupon: %v", err.Error())
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return 0, cpmodel.ErrAlreadyExists
		}
		return 0, err
	}

	return id, nil
}

func (r *dbRepo) ClaimCoupon(ctx context.Context, userID, couponName string) error {
	tx, err := r.dbCoupon.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Printf("error starting transaction: %v", err.Error())
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var (
		couponID        int64
		remainingAmount int
	)

	// lock coupon row
	err = tx.QueryRowContext(
		ctx,
		queryLockCouponByName,
		couponName,
	).Scan(&couponID, &remainingAmount)

	if err != nil {
		log.Printf("error locking coupon: %v", err.Error())
		if err == sql.ErrNoRows {
			return cpmodel.ErrNotFound
		}
		return err
	}

	if remainingAmount <= 0 {
		return cpmodel.ErrOutOfStock
	}

	// insert claim
	_, err = tx.ExecContext(
		ctx,
		queryInsertCouponClaim,
		couponID,
		userID,
	)


	if err != nil {
		log.Printf("error inserting coupon claim: %v", err.Error())
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return cpmodel.ErrAlreadyClaimed
		}
		return err
	}


	// deduct stock
	_, err = tx.ExecContext(
		ctx,
		queryUpdateCouponRemainingAmount,
		couponID,
	)

	if err != nil {
		log.Printf("error updating coupon stock: %v", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error committing transaction: %v", err.Error())
		return err
	}

	return nil
}

func (r *dbRepo) GetCouponDetailsByName(ctx context.Context, name string) (cpmodel.CouponDetail, error) {
	var result cpmodel.CouponDetail

	err := r.dbCoupon.DB.QueryRowContext(
		ctx,
		queryGetCouponDetailsByName,
		name,
	).Scan(
		&result.Name,
		&result.Amount,
		&result.RemainingAmount,
		&result.ClaimedBy,
	)

	if err != nil {
		log.Printf("error getting coupon details: %v", err)
		return cpmodel.CouponDetail{}, err
	}

	return result, nil
}
