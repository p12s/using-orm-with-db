package service

import (
	"crypto/sha1" // nolint:gosec
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/p12s/using-orm-with-db/internal/domain"
	"github.com/p12s/using-orm-with-db/internal/repository"
	"strconv"
	"time"
)

const (
	SALT               = "jk-23ds8d-fr3pwf[saf3=29f[asdf" // nolint:gosec
	HMAC_SECRET        = "fowefrewobsojh"                 // nolint:gosec
	TIME_TO_LIVE_HOURS = 24
)

var _ Auther = (*AuthService)(nil)

// Auther - authentication service interface
type Auther interface {
	CreateAccount(input domain.SignUpInput) error
	GetTokenByCredentials(input domain.SignInInput) (string, error)
	ParseToken(token string) (int, error)
	GetAccountById(id int) (domain.Account, error)
}

// AuthService - service
type AuthService struct {
	authRepo   repository.Auther
	salt       string
	hmacSecret []byte
}

// NewAuthService - constructor
func NewAuthService(authRepo repository.Auther) *AuthService {
	return &AuthService{
		authRepo:   authRepo,
		salt:       SALT,
		hmacSecret: []byte(HMAC_SECRET),
	}
}

// CreateAccount
func (s *AuthService) CreateAccount(input domain.SignUpInput) error {
	passwordHash, err := s.generatePasswordHash(input.Password)
	if err != nil {
		return fmt.Errorf("generate password hash fail: %w", err)
	}

	input.Password = passwordHash
	return s.authRepo.CreateAccount(input)
}

// GetTokenByCredentials
func (s *AuthService) GetTokenByCredentials(input domain.SignInInput) (string, error) {
	passwordHash, err := s.generatePasswordHash(input.Password)
	if err != nil {
		return "", fmt.Errorf("generate pass hash fail: %w", err)
	}

	account, err := s.authRepo.GetByCredentials(input.Email, passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user with this creds not found")
		}

		return "", fmt.Errorf("user creds wrong: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   strconv.Itoa(account.Id),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(TIME_TO_LIVE_HOURS * time.Hour).Unix(),
	})

	return token.SignedString(s.hmacSecret)
}

// ParseToken
func (s *AuthService) ParseToken(token string) (int, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.hmacSecret, nil
	})
	if err != nil {
		return 0, fmt.Errorf("token parse fail: %w", err)
	}

	if !t.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return 0, fmt.Errorf("invalid subject")
	}

	id, err := strconv.Atoi(subject)
	if err != nil {
		return 0, fmt.Errorf("invalid subject")
	}

	return id, nil
}

func (s *AuthService) GetAccountById(id int) (domain.Account, error) {
	return s.authRepo.GetAccountById(id)
}

// generatePasswordHash
func (s *AuthService) generatePasswordHash(password string) (string, error) {
	hash := sha1.New() // #nosec
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", fmt.Errorf("hash write: %w", err)
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(SALT))), nil
}
