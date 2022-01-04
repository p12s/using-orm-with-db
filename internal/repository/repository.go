package repository

import "github.com/uptrace/bun"

type Repository struct {
	Auther
}

func NewRepository(db *bun.DB) *Repository { // TODO db.DB.Conn() ??
	return &Repository{
		Auther: NewAuth(db),
	}
}
