package domain

import (
	"time"
)

type Tweets struct {
	id         string
	text       string
	likes      uint
	created_at time.Time
	updated_at time.Time
}
