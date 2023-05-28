package http

import (
	"encoding/json"
	"go-rest-api/internal/common"
	"go-rest-api/internal/tweet/adapter"
	"go-rest-api/internal/tweet/domain"
	"go-rest-api/internal/tweet/service"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type TweetRouter struct {
	Db *sqlx.DB
}

func (tr *TweetRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", tr.Create) // POST /tweet - Create a new tweet.
	return r
}

func (tr *TweetRouter) Create(w http.ResponseWriter, r *http.Request) {
	var payload domain.TweetPayload

	username, _, ok := r.BasicAuth()

	if !ok {
		status := http.StatusUnauthorized
		errorResp := common.ToErrorResponse(status)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(errorResp)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil || len(strings.TrimSpace(payload.Content)) == 0 {
		status := http.StatusBadRequest
		errorResp := common.ToErrorResponse(status)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(errorResp)
		return
	}

	payload.User_id = username

	TweetEntity := payload.ToEntity()

	uow := adapter.NewUnitOfWork(tr.Db)
	err = service.CreateTweet(uow, TweetEntity)

	if err != nil {
		status := http.StatusInternalServerError
		errorResp := common.ToErrorResponse(status)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(errorResp)
		return
	}

	status := http.StatusCreated
	errorResp := common.ToErrorResponse(status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(errorResp)
	return
}
