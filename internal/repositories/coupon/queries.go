package coupon

const (
	// create coupon
	queryInsertCoupon = `
		INSERT INTO coupons (
			cdate,
			udate,
			name,
			amount,
			remaining_amount
		) VALUES (
			NOW(),
			NOW(),
			$1,
			$2,
			$2
		) RETURNING coupon_id
		;
	`

	// lock coupon for claim (transaction)
	queryLockCouponByName = `
		SELECT
			coupon_id,
			remaining_amount
		FROM coupons
		WHERE name = $1
		FOR UPDATE
		;
	`

	// insert coupon claim
	queryInsertCouponClaim = `
		INSERT INTO coupon_claims (
			coupon_id,
			user_id,
			cdate,
			udate
		) VALUES (
			$1,
			$2,
			NOW(),
			NOW()
		)
		;
	`

	// deduct coupon stock
	queryUpdateCouponRemainingAmount = `
		UPDATE coupons
		SET
			remaining_amount = remaining_amount - 1,
			udate = NOW()
		WHERE coupon_id = $1
		;
	`

	// get coupon details
	queryGetCouponDetailsByName = `
		SELECT
			c.name,
			c.amount,
			c.remaining_amount,
			COALESCE(
				ARRAY_AGG(cc.user_id ORDER BY cc.cdate)
				FILTER (WHERE cc.user_id IS NOT NULL),
				'{}'
			) AS claimed_by
		FROM coupons c
		LEFT JOIN coupon_claims cc
			ON cc.coupon_id = c.coupon_id
		WHERE c.name = $1
		GROUP BY c.coupon_id
		;
	`
)
