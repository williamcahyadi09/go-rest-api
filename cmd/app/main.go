package main

import (
	"fmt"
	"go-rest-api/config"
	bookHttp "go-rest-api/internal/book/http"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	_ = config.GetConfig()
	con := config.InitDBConnectionPool()

	defer con.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	bookRouter := bookHttp.BookRouter{Db: con}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Mount("/book", bookRouter.Routes())

	url := "http://127.0.0.1:8000"
	fmt.Println("Listening on", url)

	http.ListenAndServe(":8000", r)
}
