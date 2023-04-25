package domain

import (
	"time"
)

type Tweets struct {
	id         string
	content       string
	likes      uint
	user_id    string
	created_at time.Time
	updated_at time.Time
}
