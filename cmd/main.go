package main

import (
	"log"
	"net/http"

	"github.com/ktcf/hello-api/config"
	"github.com/ktcf/hello-api/handlers"
	"github.com/ktcf/hello-api/handlers/rest"
	"github.com/ktcf/hello-api/translation"
)

func main() {
	cfg := config.LoadConfiguration()
	addr := cfg.Port

	mux := http.NewServeMux()

	var translationService rest.Translator
	translationService = translation.NewStaticService()

	if cfg.LegacyEndpoint != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoint)
		client := translation.NewHelloClient(cfg.LegacyEndpoint)
		translationService = translation.NewRemoteService(client)
	}

	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/info", handlers.Info)

	log.Printf("listening on %s", addr)

	log.Fatal(http.ListenAndServe(addr, mux)) // nolint
}
