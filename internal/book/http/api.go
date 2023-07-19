package http

import (
	"encoding/json"
	"go-rest-api/internal/book/adapter"
	"go-rest-api/internal/book/domain"
	"go-rest-api/internal/book/http/schema"
	"go-rest-api/internal/book/service"
	"go-rest-api/internal/common"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type BookRouter struct {
	Db *sqlx.DB
}

func (br *BookRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", br.Create)     // POST /book - Create a new book.
	r.Get("/{id}", br.GetById) // POST /book/{id} - Get book by id.
	return r
}

func (br *BookRouter) Create(w http.ResponseWriter, r *http.Request) {
	var payload domain.BookPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		status := http.StatusBadRequest
		errorResp := common.ToErrorResponse(status)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(errorResp)
		return
	}

	BookEntity := domain.NewBook(
		payload.Title,
		payload.Author,
		payload.Description,
		payload.Price,
	)

	uow := adapter.NewUnitOfWork(br.Db)
	err = service.CreateBook(uow, BookEntity)

	if err != nil {
		status := http.StatusInternalServerError
		errorResp := common.ToErrorResponse(status)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(errorResp)
		return
	}

	common.RespondwithJSON(
		w,
		http.StatusCreated,
		schema.NewBookResponse(
			"Create book successful",
			BookEntity,
		),
	)
	return
}

func (br *BookRouter) GetById(w http.ResponseWriter, r *http.Request) {
	book_id := chi.URLParam(r, "id")
	uow := adapter.NewUnitOfWork(br.Db)
	book, err := service.GetBookById(uow, book_id)

	if err != nil {
		status := http.StatusInternalServerError
		errorResp := common.ToErrorResponse(status)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(errorResp)
		return
	}

	common.RespondwithJSON(
		w,
		http.StatusOK,
		schema.NewBookResponse(
			"Get book successful",
			book,
		),
	)
	return
}
