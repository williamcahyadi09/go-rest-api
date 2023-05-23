package domain

import (
	"github.com/jmoiron/sqlx"
)

type TweetRepositoryInterface interface {
	GetByUserId(tx *sqlx.Tx, user_id string) ([]Tweet, error)
	CreateTweet(tx *sqlx.Tx, tweet *Tweet) error
	// Update(ctx context.Context, tweet *Tweet) error
}

type UnitOfWorkInterface interface {
	GetTweetRepo() TweetRepositoryInterface
	Begin() (*sqlx.Tx, error)
	Commit(tx *sqlx.Tx) error
	Rollback(tx *sqlx.Tx) error
}
