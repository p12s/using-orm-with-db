package service

import (
	"github.com/p12s/using-orm-with-db/internal/repository"
)

// Service - just service
type Service struct {
	Auther
}

// NewService - constructor
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auther: NewAuthService(repos.Auther),
	}
}
