package app

import (
	"log"
	"time"

	v1 "coupon-system/internal/controller/http/v1"
	grace "coupon-system/internal/grace"

	"coupon-system/internal/httpclient"

	"coupon-system/internal/postgres"

	"coupon-system/config"

	// common

	// services
	cpsvc "coupon-system/internal/services/coupon"
	pingsvc "coupon-system/internal/services/ping"

	// repositories
	cprepo "coupon-system/internal/repositories/coupon"
	pingrepo "coupon-system/internal/repositories/ping"
)

type appHttp struct {
	v1Controller v1.V1Controller
}

// RegisterHandlers registers the http handlers
func NewAppHTTP(
	config *config.Config,
) *appHttp {
	// postgres
	database, err := postgres.New(&postgres.Config{
		ServiceName:   config.Database.ServiceName,
		Dsn:           config.Database.DSN,
		MaxConn:       config.Database.MaxOpenConn,
		MaxIdle:       config.Database.MaxIdleConn,
	})
	if err != nil {
		log.Fatalf("error init postgres %s", err.Error())
	}

	// init repositories
	pingRepo := pingrepo.NewPingRepository(database)
	couponRepo := cprepo.NewCouponRepository(database)

	
	// init repositories
	pingService := pingsvc.NewPingService(pingRepo)
	couponService := cpsvc.NewCouponService(couponRepo)
	
	// init controllers
	v1Controller := v1.NewV1Controller(
		pingService,
		couponService,
	)


	return &appHttp{
		v1Controller: v1Controller,
	}
}

// Run runs the http app
func (a *appHttp) Run(config *config.Config) {
	log.Printf("HTTP server running on port %s\n", config.Port)
	// run http server
	grace.Serve(
		config.Port,
		a.RegisterHandlers(config),
	)
}

func SetupHttpClient(cfg *config.Config) httpclient.Client {
	httpClientCfg := &httpclient.Config{
		Timeout: cfg.HTTPClient.TimeoutMS,
		Transport: struct {
			DisableKeepAlives   bool
			MaxIdleConns        int
			MaxConnsPerHost     int
			MaxIdleConnsPerHost int
			IdleConnTimeout     time.Duration
		}{
			DisableKeepAlives:   cfg.HTTPClient.DisableKeepAlives,
			MaxIdleConns:        cfg.HTTPClient.MaxIdleConns,
			MaxConnsPerHost:     cfg.HTTPClient.MaxConnsPerHost,
			MaxIdleConnsPerHost: cfg.HTTPClient.MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(cfg.HTTPClient.IdleConnTimeout) * time.Second,
		},
	}

	return httpclient.New(httpClientCfg)
}
