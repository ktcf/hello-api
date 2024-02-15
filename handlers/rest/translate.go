package rest

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Translator interface {
	Translate(word string, language string) string
}

// TranslateHandler will translate calls for caller.
type TranslateHandler struct {
	service Translator
}

// NewTranslateHandler will create a new instance if the handler using a
// translation service.
func NewTranslateHandler(service Translator) *TranslateHandler {
	return &TranslateHandler{service}
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

// TranslateHandler will take a given request with a path value of the
// word to be translated and a query parameter of the
// language to translate to.
func (t *TranslateHandler) TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translate := t.service.Translate(word, language)
	if translate == "" {
		w.WriteHeader(404)
		return
	}
	resp := Resp{
		Language:    language,
		Translation: translate,
	}
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
