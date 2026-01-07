-- drop indexes on coupon_claims
DROP INDEX IF EXISTS idx_coupon_claims_user_id;
DROP INDEX IF EXISTS idx_coupon_claims_coupon_id;

-- drop coupon_claims table and sequence
DROP TABLE IF EXISTS coupon_claims;
DROP SEQUENCE IF EXISTS coupon_claims_id_seq;



-- drop indexes on coupons
DROP INDEX IF EXISTS idx_coupons_name;

-- drop coupons table and sequence
DROP TABLE IF EXISTS coupons;
DROP SEQUENCE IF EXISTS coupons_id_seq;
