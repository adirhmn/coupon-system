-- sequence for coupons
CREATE SEQUENCE IF NOT EXISTS coupons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- coupons table
CREATE TABLE IF NOT EXISTS coupons (
    cdate TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    udate TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    coupon_id BIGINT NOT NULL DEFAULT nextval('coupons_id_seq'::regclass),

    name VARCHAR(100) NOT NULL,
    amount INT NOT NULL,
    remaining_amount INT NOT NULL,

    CHECK (amount >= 0),
    CHECK (remaining_amount >= 0),

    PRIMARY KEY (coupon_id),
    UNIQUE (name)
);

-- index for coupon name lookup
CREATE INDEX IF NOT EXISTS idx_coupons_name
ON coupons (name);



-- sequence for coupon_claims
CREATE SEQUENCE IF NOT EXISTS coupon_claims_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- coupon_claims table
CREATE TABLE IF NOT EXISTS coupon_claims (
    cdate TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    udate TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    coupon_claim_id BIGINT NOT NULL DEFAULT nextval('coupon_claims_id_seq'::regclass),

    coupon_id BIGINT NOT NULL,
    user_id VARCHAR(100) NOT NULL,

    PRIMARY KEY (coupon_claim_id),

    CONSTRAINT fk_coupon_claims_coupon
        FOREIGN KEY (coupon_id)
        REFERENCES coupons (coupon_id)
        ON DELETE CASCADE,

    CONSTRAINT uq_coupon_claims_coupon_user
        UNIQUE (coupon_id, user_id)
);

-- index for coupon_claims
CREATE INDEX IF NOT EXISTS idx_coupon_claims_coupon_id
ON coupon_claims (coupon_id);

CREATE INDEX IF NOT EXISTS idx_coupon_claims_user_id
ON coupon_claims (user_id);
