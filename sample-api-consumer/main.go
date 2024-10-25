package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Setup service
	r.Post("/api/v1/message", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(string(body))
		w.WriteHeader(200)
		w.Write([]byte("Success"))
	})

	fmt.Println("Service started on port 3000")

	http.ListenAndServe(":3001", r)

}
