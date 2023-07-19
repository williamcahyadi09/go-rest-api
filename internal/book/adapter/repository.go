package adapter

import (
	"database/sql"
	"go-rest-api/internal/book/domain"

	"github.com/jmoiron/sqlx"
)

type BookSQL struct {
	Id          string       `db:"id"`
	Title       string       `db:"title"`
	Author      string       `db:"author"`
	Description string       `db:"description"`
	Price       uint         `db:"price"`
	Created_at  sql.NullTime `db:"created_at"`
	Updated_at  sql.NullTime `db:"updated_at"`
	Deleted_at  sql.NullTime `db:"deleted_at"`
}

func EntityToSql(book *domain.Book) BookSQL {
	return BookSQL{
		Id:          book.Id,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		Price:       book.Price,
		Created_at:  book.Created_at,
		Updated_at:  book.Updated_at,
		Deleted_at:  book.Deleted_at,
	}
}

func (book *BookSQL) ToEntity() domain.Book {
	return domain.Book{
		Id:          book.Id,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		Price:       book.Price,
		Created_at:  book.Created_at,
		Updated_at:  book.Updated_at,
		Deleted_at:  book.Deleted_at,
	}
}

type BookRepository struct{}

func NewBookRepository() domain.BookRepositoryInterface {
	return &BookRepository{}
}

func (r *BookRepository) GetById(tx *sqlx.Tx, id string) (*domain.Book, error) {
	var book domain.Book
	query := `SELECT * FROM book WHERE id=$1 ORDER BY created_at LIMIT 1`
	err := tx.QueryRowx(query, id).StructScan(&book)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) CreateBook(tx *sqlx.Tx, book *domain.Book) error {
	query := `INSERT INTO book (id, title, author, description, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	bookSql := EntityToSql(book)

	_, err := tx.Exec(
		query,
		bookSql.Id,
		bookSql.Title,
		bookSql.Author,
		bookSql.Description,
		bookSql.Price,
		bookSql.Created_at,
		bookSql.Updated_at,
	)

	if err != nil {
		return err
	}
	return nil
}
