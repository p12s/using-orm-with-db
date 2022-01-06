package repository

import (
	"context"
	"fmt"
	"github.com/p12s/using-orm-with-db/internal/domain"
	"github.com/uptrace/bun"
	"time"
)

var _ Auther = (*Auth)(nil)

// Auther - auth repository interface
type Auther interface {
	CreateAccount(account domain.SignUpInput) error
	GetAccountById(id int) (domain.Account, error)
	GetByCredentials(email, password string) (domain.Account, error)
}

// Auth
type Auth struct {
	db *bun.DB
}

// Auth - constructor
func NewAuth(db *bun.DB) *Auth {
	return &Auth{db: db}
}

// CreateAccount
func (r *Auth) CreateAccount(input domain.SignUpInput) error {
	t := time.Now()
	account := &domain.Account{
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: &t,
	}

	_, err := r.db.NewInsert().Model(account).Exec(context.TODO())
	if err != nil {
		return fmt.Errorf("repo create account fail: %w", err)
	}
	return nil
}

// GetAccount
func (r *Auth) GetAccountById(id int) (domain.Account, error) {
	account := &domain.Account{}
	err := r.db.NewSelect().Model(account).
		Where("id=?", id).
		Scan(context.TODO())
	if err != nil {
		return *account, fmt.Errorf("repo get-by-id account fail: %w", err)
	}

	return *account, nil
}

// GetByCredentials
func (r *Auth) GetByCredentials(email, password string) (domain.Account, error) {
	account := &domain.Account{}
	err := r.db.NewSelect().Model(account).
		Where("email=?", email).
		Where("password_hash=?", password).
		Scan(context.TODO())
	if err != nil {
		return *account, fmt.Errorf("repo get-by-creds account fail: %w", err)
	}

	return *account, nil
}
