package service

import (
	"go-rest-api/internal/tweets/domain"
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
		return err
	}
	return nil
}
