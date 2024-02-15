package main

import (
	"github.com/ktcf/hello-api/handlers"
	"github.com/ktcf/hello-api/handlers/rest"
	"github.com/ktcf/hello-api/translation"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s", addr)

	log.Fatal(http.ListenAndServe(addr, mux)) // nolint
}
