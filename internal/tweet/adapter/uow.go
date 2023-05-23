package adapter

import (
	"fmt"
	"go-rest-api/internal/tweet/domain"

	"github.com/jmoiron/sqlx"
)

type UnitOfWork struct {
	Db              *sqlx.DB
	TweetRepository domain.TweetRepositoryInterface
}

func NewUnitOfWork(
	db *sqlx.DB,
) domain.UnitOfWorkInterface {
	return &UnitOfWork{
		Db:              db,
		TweetRepository: NewTweetRepository(),
	}
}

func (uow *UnitOfWork) Begin() (*sqlx.Tx, error) {
	tx, err := uow.Db.Beginx()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (uow *UnitOfWork) Commit(tx *sqlx.Tx) error {
	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("error when committing transaction: %v", err)
	}
	return nil
}

func (uow *UnitOfWork) Rollback(tx *sqlx.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return fmt.Errorf("error when rollback transaction: %v", err)
	}
	return nil
}

func (uow *UnitOfWork) GetTweetRepo() domain.TweetRepositoryInterface {
	return uow.TweetRepository
}
