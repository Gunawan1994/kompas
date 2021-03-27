package main

import (
	"net/http"
	"time"

	"kompas/pkg/grace"
)

func startServer(handler http.Handler, port string) error {

	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler,
	}

	return grace.Serve(port, srv)
}
