package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver

	"kompas/internal/config"

	logHdl "kompas/internal/delivery/http/login"

	thnHdl "kompas/internal/delivery/http/dashboard"

	thnRes "kompas/internal/resource/dashboard"

	usrRes "kompas/internal/resource/user"

	logSvc "kompas/internal/service/login"

	thnSvc "kompas/internal/service/dashboard"

	audRes "kompas/internal/resource/log"

	"kompas/pkg/cache"
)

func startApp() error {
	config.Init()
	cfg := config.Get()

	db, err := sqlx.Open(cfg.DB.Driver, cfg.DB.Master)
	if err != nil {
		return err
	}

	redis := cache.NewRedis(cache.RedisConfig{
		Address: cfg.Cache.Address,
	})

	audRes := audRes.New(db)

	usrRes := usrRes.New(db, redis)
	logSvc := logSvc.New(usrRes)
	logHdl := logHdl.New(logSvc)

	thnRes := thnRes.New(db)
	thnSvc := thnSvc.New(thnRes, audRes)
	thnHdl := thnHdl.New(thnSvc)

	r := newRouter(
		logHdl, thnHdl,
	)
	return startServer(r, cfg.Server.HTTPPort)
}
