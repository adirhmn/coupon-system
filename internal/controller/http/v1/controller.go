package v1

import (
	"github.com/gin-gonic/gin"

	// services
	cpsvc "coupon-system/internal/services/coupon"
	pingsvc "coupon-system/internal/services/ping"
)

type V1Controller interface {
	Ping(c *gin.Context)
	CreateCoupon(c *gin.Context)
	ClaimCoupon(c *gin.Context)
	DetailCoupon(c *gin.Context)
}

type v1Controller struct {
	pingService                   pingsvc.PingServiceProvider
	couponService cpsvc.CouponServiceProvider
}

func NewV1Controller(
	pingService pingsvc.PingServiceProvider,
	couponService cpsvc.CouponServiceProvider,
	) V1Controller {
	return &v1Controller{
		pingService:                   pingService,
		couponService: couponService,
	}
}