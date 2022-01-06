package repository

import "github.com/uptrace/bun"

//go:generate mockgen -destination mocks/mock.go -package repository github.com/p12s/using-orm-with-db/internal/repository Auther

type Repository struct {
	Auther
}

func NewRepository(db *bun.DB) *Repository { // TODO db.DB.Conn() ??
	return &Repository{
		Auther: NewAuth(db),
	}
}
