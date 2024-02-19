package handlers

import (
	"encoding/json"
	"net/http"
)

var (
	tag  string
	hash string
	date string
)

func Info(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{
		"tag":  tag,
		"hash": hash,
		"date": date,
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
