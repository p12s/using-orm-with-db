package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/uptrace/bun"
	"time"
)

var validate *validator.Validate // nolint:gochecknoglobals

// init - init validator
func init() {
	validate = validator.New()
}

type Role int

const (
	ROLE_CUSTOMER Role = iota
)

// Account
type Account struct {
	bun.BaseModel `bun:"table:account,alias:a"`

	Id        int        `json:"id,omitempty" bun:"id,pk,autoincrement"`
	Email     string     `json:"email" bun:"email" binding:"required"`
	Password  string     `json:"-" bun:"password_hash" binding:"required"`
	Role      Role       `json:"role" bun:"role"`
	CreatedAt *time.Time `json:"created_at,omitempty" bun:"created_at"` // nolint
}

// SignUpInput
type SignUpInput struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password,omitempty" binding:"required" validate:"required,gte=6"`
}

func (i *SignUpInput) Validate() error {
	return validate.Struct(i)
}

// SignInInput
type SignInInput struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,gte=6"`
}

func (i *SignInInput) Validate() error {
	return validate.Struct(i)
}
