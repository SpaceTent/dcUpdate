package app

import (
	l "log/slog"
	"net/http"

	"dcupdate/app/controllers"
)

func Start() {

	l.Info("Starting")

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/update", controllers.Update)

	http.ListenAndServe(":8080", nil)
}
