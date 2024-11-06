// Package service contains business logic handler
package service

import "github.com/ak1m1tsu/urler/internal/pkg/config"

// Service represents a business logic handler
type Service struct {
	cfg *config.Config
}

// New creates new business logic hander
func New(cfg *config.Config) *Service {
	return &Service{cfg: cfg}
}
