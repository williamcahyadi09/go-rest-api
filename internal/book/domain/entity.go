package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id          string
	Title       string
	Author      string
	Description string
	Price       uint
	Created_at  sql.NullTime
	Updated_at  sql.NullTime
	Deleted_at  sql.NullTime
}

func NewBook(
	title string,
	author string,
	description string,
	price uint,
) *Book {
	sqlTime := sql.NullTime{
		time.Now().UTC(),
		true,
	}
	return &Book{
		Id:          uuid.New().String(),
		Title:       title,
		Author:      author,
		Description: description,
		Price:       price,
		Created_at:  sqlTime,
		Updated_at:  sqlTime,
	}
}

type BookPayload struct {
	Title       string
	Author      string
	Description string
	Price       uint
}
