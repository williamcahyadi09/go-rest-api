package domain

import (
	"time"
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
