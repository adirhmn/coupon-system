package app

import (
	"coupon-system/config"

	"github.com/gin-gonic/gin"
)

// RegisterHandler registers all the handlers
func (a *appHttp) RegisterHandlers(config *config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	// v1 group
	a.registerV1(handler.Group("api"), config)

	return handler
}

func (a *appHttp) registerV1(r *gin.RouterGroup, cfg *config.Config) {
	r.GET("/ping", a.v1Controller.Ping)
	r.GET("/coupons/:coupon_name", a.v1Controller.DetailCoupon)
	r.POST("/coupons", a.v1Controller.CreateCoupon)
	r.POST("/coupons/claim", a.v1Controller.ClaimCoupon)
}