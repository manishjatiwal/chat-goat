package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RunHTTPServer() {
    r := chi.NewRouter()

    r.Get("/", func (w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello Goat!"))
    })

    http.ListenAndServe(":3000", r)
}
