package main

import (
	"github.com/gorilla/mux"

	"kompas/cmd/http-server/middleware"

	"kompas/internal/delivery/http/login"

	dashboard "kompas/internal/delivery/http/dashboard"
)

func newRouter(
	loginHandler login.Handler,
	dashboardHandler dashboard.Handler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler.Login).Methods("POST")

	r.HandleFunc("/artikel/{id}", middleware.TokenAuthMiddleware(dashboardHandler.Get)).Methods("GET")
	r.HandleFunc("/artikel", middleware.TokenAuthMiddleware(dashboardHandler.Post)).Methods("POST")
	r.HandleFunc("/artikel", middleware.TokenAuthMiddleware(dashboardHandler.GetAll)).Methods("GET")

	return r
}
