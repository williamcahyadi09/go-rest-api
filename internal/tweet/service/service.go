package service

import (
	"go-rest-api/internal/tweet/domain"
)

func CreateTweet(
	uow domain.UnitOfWorkInterface,
	tweet *domain.Tweet,
) error {
	tx, err := uow.Begin()

	if err != nil {
		return err
	}

	err = uow.GetTweetRepo().CreateTweet(tx, tweet)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// func GetTweetsByUserId(user_id string) {

// }
