package faas

import (
	"github.com/ktcf/hello-api/handlers/rest"
	"github.com/ktcf/hello-api/translation"
	"net/http"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)

	translateHandler.TranslateHandler(w, r)
}
