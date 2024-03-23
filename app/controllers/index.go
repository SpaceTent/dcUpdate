package controllers

import (
	l "log/slog"
	"net/http"

	"dcupdate/pkg/web"
)

func Index(w http.ResponseWriter, r *http.Request) {

	l.Debug("Index")

	AC, err := Get(r, "Index")
	if err != nil {
		l.With("Error", err).Error("Error getting AppContext")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	web.Render(w, r, AC.PageData, "index.gohtml")

}
