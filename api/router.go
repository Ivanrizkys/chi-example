package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// import "net/http"
type Body struct {
	Name string `json:"nama"`
}

func Router(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		// * get query params in chi
		params := r.URL.Query().Get("role")
		// * get request body
		decoder := json.NewDecoder(r.Body)
		var t Body
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(params)
		fmt.Println(t.Name)
		w.Write([]byte(t.Name))
	})
}
