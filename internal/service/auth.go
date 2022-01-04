package service

import (
	"github.com/p12s/using-orm-with-db/internal/repository"
)

var _ Auther = (*AuthService)(nil)

// Auther - authentication service interface
type Auther interface {
}

// AuthService - service
type AuthService struct {
	authRepo repository.Auther
}

// NewAuthService - constructor
func NewAuthService(authRepo repository.Auther) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}
