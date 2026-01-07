package coupon

import "errors"

var (
	ErrAlreadyClaimed = errors.New("coupon already claimed")
	ErrOutOfStock     = errors.New("coupon out of stock")
	ErrNotFound       = errors.New("coupon not found")
	ErrAlreadyExists  = errors.New("coupon already exists")
)