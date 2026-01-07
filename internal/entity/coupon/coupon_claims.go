package coupon

import "time"

type CouponClaim struct {
	ID       int64     `json:"id"`
	CouponID int64     `json:"coupon_id"`
	UserID   string    `json:"user_id"`
	CDate    time.Time `json:"cdate"`
	UDate    time.Time `json:"udate"`
}
