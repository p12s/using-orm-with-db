package service

import (
	"github.com/p12s/using-orm-with-db/internal/repository"
)

//go:generate mockgen -destination mocks/mock.go -package service github.com/p12s/using-orm-with-db/internal/service Auther

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
