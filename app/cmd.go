package app

import (
	l "log/slog"
	"net/http"

	"dcupdate/app/controllers"
	"dcupdate/app/services/docker"
)

func Start() {

	l.Info("Starting")

	// Need the link to the Docker Compose File

	_, _ = docker.GetRunningContainers()

	// docker.GetNewImage(c[0], "v0.0.144")

	// docker.RestartContainer(c[0], "v0.0.144")

	// docker.RestartContainer(c[0], "v0.0.146")

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/update", controllers.Update)

	http.ListenAndServe(":8080", nil)
}
