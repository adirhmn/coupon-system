package coupon

import "time"

type Coupon struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Amount          int       `json:"amount"`
	RemainingAmount int       `json:"remaining_amount"`
	CDate           time.Time `json:"cdate"`
	UDate           time.Time `json:"udate"`
}
