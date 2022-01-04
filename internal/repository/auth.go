package repository

import (
	"github.com/uptrace/bun"
)

var _ Auther = (*Auth)(nil)

// Auther - auth repository interface
type Auther interface {
}

// Auth
type Auth struct {
	db *bun.DB // TODO db.DB.Conn() ??
}

// Auth - constructor
func NewAuth(db *bun.DB) *Auth {
	return &Auth{db: db}
}
