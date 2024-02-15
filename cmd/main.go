package main

import (
	"github.com/ktcf/hello-api/handlers"
	"github.com/ktcf/hello-api/handlers/rest"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s", addr)

	log.Fatal(http.ListenAndServe(addr, mux)) // nolint
}
