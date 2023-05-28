package domain

import (
	"time"

	"github.com/google/uuid"
)

type Tweet struct {
	Id         string
	Content    string
	Likes      uint
	User_id    string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type TweetPayload struct {
	Content string
	User_id string
}

func (payload *TweetPayload) ToEntity() *Tweet {
	time_created := time.Now().UTC()
	return &Tweet{
		Id:         uuid.New().String(),
		Content:    payload.Content,
		Likes:      0,
		User_id:    payload.User_id,
		Created_at: time_created,
		Updated_at: time_created,
	}
}
