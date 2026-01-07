package v1

import (
	"context"
	serverctrl "coupon-system/internal/controller/http"
	cpmodel "coupon-system/internal/entity/coupon"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (v1 *v1Controller) CreateCoupon(c *gin.Context) {
	ctx, cancelCtx := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancelCtx()

	var req struct {
		Name     string `json:"name" binding:"required"`
		Amount int    `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		serverctrl.ResponseHandler(c, http.StatusBadRequest, nil, fmt.Errorf("invalid input"))
		return
	}

	_, err := v1.couponService.InsertCoupon(ctx, req.Name, req.Amount)
	if err != nil {
		serverctrl.ResponseHandler(c, getErrorHTTPStatus(err), nil, err)
		return
	}

	serverctrl.ResponseHandler(c, http.StatusCreated, "coupon created", nil)
}

func (v1 *v1Controller) ClaimCoupon(c *gin.Context) {
	ctx, cancelCtx := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancelCtx()

	var req struct {
		UserID     string `json:"user_id" binding:"required"`
		CouponName string    `json:"coupon_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		serverctrl.ResponseHandler(c, http.StatusBadRequest, nil, fmt.Errorf("invalid input"))
		return
	}

	err := v1.couponService.ClaimCoupon(ctx, req.UserID, req.CouponName)
	if err != nil {
		serverctrl.ResponseHandler(c, getErrorHTTPStatus(err), nil, err)
		return
	}

	serverctrl.ResponseHandler(c, http.StatusOK, "coupon claimed", nil)
}


func (v1 *v1Controller) DetailCoupon(c *gin.Context) {
	ctx, cancelCtx := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancelCtx()

	couponName := c.Param("coupon_name")
	if couponName == ""{
		serverctrl.ResponseHandler(c, http.StatusBadRequest, nil, fmt.Errorf("invalid input"))
		return
	}

	couponDetails, err := v1.couponService.GetCouponDetailsByName(ctx, couponName)
	if err != nil {
		serverctrl.ResponseHandler(c, getErrorHTTPStatus(err), nil, err)
		return
	}

	c.JSON(http.StatusOK, couponDetails)
	return
}


func getErrorHTTPStatus( err error) int {
	switch err {
		case cpmodel.ErrAlreadyExists:
			return http.StatusConflict
		case cpmodel.ErrOutOfStock:
			return http.StatusBadRequest
		case cpmodel.ErrNotFound:
			return http.StatusNotFound
		case cpmodel.ErrAlreadyClaimed:
			return http.StatusConflict
		default:
			return http.StatusInternalServerError
	}
}