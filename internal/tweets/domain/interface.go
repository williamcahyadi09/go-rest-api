package domain

type CustomerRepository interface {
	Get(id string) (*Tweets, error)
	Add(tweet Tweets) error
	Update(tweet Tweets) error
}
