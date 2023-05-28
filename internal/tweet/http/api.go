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
		errorResp, _ := common.ToErrorResponse(common.Unauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(errorResp)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil || len(strings.TrimSpace(payload.Content)) == 0 {
		errorResp, parseError := common.ToErrorResponse(common.BadRequest)

		if parseError != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorResp)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorResp)
		return
	}

	payload.User_id = username

	TweetEntity := payload.ToEntity()

	uow := adapter.NewUnitOfWork(tr.Db)
	err = service.CreateTweet(uow, TweetEntity)

	if err != nil {
		errorResp, _ := common.ToErrorResponse(common.Unauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorResp)
		return
	}

	return
}
