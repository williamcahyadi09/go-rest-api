package adapter

import (
	"go-rest-api/internal/tweets/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type TweetSQL struct {
	Id         string    `db:"id"`
	Content    string    `db:"content"`
	Likes      uint      `db:"likes"`
	User_id    string    `db:"user_id"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
	Deleted_at time.Time `db:"deleted_at"`
}

func EntityToSql(t *domain.Tweet) TweetSQL {
	return TweetSQL{
		Id:         t.Id,
		Content:    t.Content,
		Likes:      t.Likes,
		User_id:    t.User_id,
		Created_at: t.Created_at,
		Updated_at: t.Updated_at,
		Deleted_at: t.Deleted_at,
	}
}

func (t *TweetSQL) ToEntity() domain.Tweet {
	return domain.Tweet{
		Id:         t.Id,
		Content:    t.Content,
		Likes:      t.Likes,
		User_id:    t.User_id,
		Created_at: t.Created_at,
		Updated_at: t.Updated_at,
		Deleted_at: t.Deleted_at,
	}
}

type TweetRepository struct{}

func NewTweetRepository() domain.TweetRepositoryInterface {
	return &TweetRepository{}
}

func (r *TweetRepository) GetByUserId(tx *sqlx.Tx, user_id string) ([]domain.Tweet, error) {
	query := `SELECT * FROM tweets WHERE user_id=$1`
	rows, err := tx.Queryx(query, user_id)

	if err != nil {
		return nil, err
	}

	tweets := []domain.Tweet{}
	for rows.Next() {
		var tweetSql TweetSQL
		if err := rows.StructScan(&tweetSql); err != nil {
			return nil, err
		}
		tweet := tweetSql.ToEntity()
		tweets = append(tweets, tweet)
	}

	// Don't know if I need this code or not
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tweets, nil
}

func (r *TweetRepository) CreateTweet(tx *sqlx.Tx, tweet *domain.Tweet) error {
	query := `INSERT INTO tweets (id, user_id, content, likes, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	tweetSql := EntityToSql(tweet)

	_, err := tx.Exec(query, tweetSql.Id, tweetSql.User_id, tweetSql.Content, tweetSql.Likes, tweetSql.Created_at, tweetSql.Updated_at)

	if err != nil {
		return err
	}
	return nil
}
