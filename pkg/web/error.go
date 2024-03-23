package web

import (
	l "log/slog"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request, code int, message string, e error) {
	l.With("error", e).Error(message)
	WriteTEXT(w, code, message)
}
