package app

import (
	l "log/slog"
	"net/http"

	"dcupdate/app/controllers"
	"dcupdate/app/environment"
)

func Start() {

	l.Info("Starting")

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/update", controllers.Update)

	http.ListenAndServe(":"+environment.GetEnvString("BIND", "9000"), nil)
}
