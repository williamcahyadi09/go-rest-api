package domain

import (
	"github.com/jmoiron/sqlx"
)

type BookRepositoryInterface interface {
	GetById(tx *sqlx.Tx, id string) (*Book, error)
	CreateBook(tx *sqlx.Tx, tweet *Book) error
	// UpdateBook(tx *sqlx.Tx, tweet *Book) error
}

type UnitOfWorkInterface interface {
	GetBookRepo() BookRepositoryInterface
	Begin() (*sqlx.Tx, error)
	Commit(tx *sqlx.Tx) error
	Rollback(tx *sqlx.Tx) error
}
