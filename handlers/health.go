// Package handlers houses other handlers needed for the API
package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{"status": "ok"}
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode json response")
	}
}
