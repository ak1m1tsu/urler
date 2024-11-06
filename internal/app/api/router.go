// Package api contains HTTP router with API endpoints
package api

import (
	"github.com/ak1m1tsu/urler/internal/pkg/service"
	"github.com/go-chi/chi/v5"
)

// Router creates configured API router
func Router(svc *service.Service) *chi.Mux {
	router := chi.NewRouter()

	return router
}
