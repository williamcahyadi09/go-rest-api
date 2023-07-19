package schema

import (
	"go-rest-api/internal/book/domain"
	"time"
)

type BookResponseModel struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Price       uint      `json:"price"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type BookResponse struct {
	Message string             `json:"message"`
	Book    *BookResponseModel `json:"book"`
}

func NewBookResponse(
	message string,
	bookEntity *domain.Book,
) *BookResponse {
	return &BookResponse{
		Message: message,
		Book: &BookResponseModel{
			Id:          bookEntity.Id,
			Title:       bookEntity.Title,
			Author:      bookEntity.Author,
			Description: bookEntity.Description,
			Price:       bookEntity.Price,
			Created_at:  bookEntity.Created_at.Time,
			Updated_at:  bookEntity.Updated_at.Time,
		},
	}
}
