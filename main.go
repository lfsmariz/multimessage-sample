package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/multimessage-sample/client"
	"github.com/multimessage-sample/producer"
)

func main() {
	service := "kafka"

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Setup service
	r.Post("/setup-service", func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Service string `json:"service"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		service = requestBody.Service
		w.WriteHeader(200)
		w.Write([]byte("Success"))
	})

	r.Get("/verify-service", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(service))
	})

	r.Post("/send-message", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if service == "kafka" {
			producer.PublishMessageKafka(string(body))
		} else if service == "rabbitmq" {
			producer.PublishMessageRabbitmq(string(body))
		}

		w.WriteHeader(200)
		w.Write([]byte("Success"))
	})

	client.StartConsumers()

	fmt.Println("Service started on port 3000")

	http.ListenAndServe(":3000", r)

}
