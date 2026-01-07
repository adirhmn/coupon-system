package coupon

import "github.com/lib/pq"

type CouponDetail struct {
	Name            string         `json:"name"`
	Amount          int            `json:"amount"`
	RemainingAmount int            `json:"remaining_amount"`
	ClaimedBy       pq.StringArray `json:"claimed_by"`
}
