package controllers

import (
	l "log/slog"
	"net/http"

	"dcupdate/app/services/DockerAPI"
	"dcupdate/app/services/dockerCompose"
	"dcupdate/pkg/web"
)

func Update(w http.ResponseWriter, r *http.Request) {

	l.Debug("Index")

	_, err := Get(r, "Index")
	if err != nil {
		l.With("Error", err).Error("Error getting AppContext")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	Image := web.GetQueryString(r, "Image")
	Tag := web.GetQueryString(r, "Tag")

	l.Info("Update Image: " + Image + " Tag: " + Tag)

	out := DockerAPI.DCContainer{}
	RunningContainers, _ := DockerAPI.GetRunningContainers()
	for _, c := range RunningContainers {
		if c.Name == Image {
			out = c
			if Tag != "" {
				go func() {

					// Update the Tags in the Compose File
					dockerCompose.UpdateComposer(c, Tag)

					// Tell Docker to Download any new Images
					DockerAPI.GetNewImage(c, Tag)

					// We need to restart the container.  But with the new Tag Details..
					// Ideally I'd like to Kill the Container and recreate it with the Same configuration
					// But that's a bit complex for now,  so I'm hacking it ...

					// docker.RestartContainer(c)
					// docker.StopContainer(c)
					// docker.CreateContainer(c, Tag)

					// Restart with a Docker Compose Command
					dockerCompose.Restart()

				}()
			}
		}
	}

	if out.Name == "" {
		web.WriteJSON(w, http.StatusNotFound, map[string]string{"Error": "Container Not Found"})
		return
	}

	web.WriteJSON(w, http.StatusOK, out)
}
