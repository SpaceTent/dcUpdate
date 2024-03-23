package web

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteHTML(w http.ResponseWriter, status int, v string) {
	w.Header().Add("Content-Type", "text/html")
	_write(w, status, v)
}

func WriteTEXT(w http.ResponseWriter, status int, v string) {
	w.Header().Add("Content-Type", "text/plain")
	_write(w, status, v)
}

func _write(w http.ResponseWriter, status int, v string) {
	w.WriteHeader(status)

	_, _ = w.Write([]byte(v))
}
