package dockerCompose

import (
	l "log/slog"
	"os"
	"strings"

	"dcupdate/app/environment"
	"dcupdate/app/services/docker"
)

func UpdateComposer(localContainer docker.DCContainer, NewTag string) {

	// Here is where it's a bit hacky, we are going to update the container with a new Tag
	// so we update the Docker Compose File, and then issue a restart

	// This is a Massive Hack, but it works for now.

	// Read in the Compose File (docker-compose.yml)
	content, err := os.ReadFile(environment.GetEnvString("COMPOSE_FILE", "local-test-data/docker-compose.yml"))
	if err != nil {
		l.With("Error", err).Error("Error reading docker-compose.yml")
		return
	}

	// Update the Tag for the Container
	NewContent := strings.ReplaceAll(string(content), localContainer.Image, localContainer.Name+":"+NewTag)

	// Write out the Compose File
	err = os.WriteFile(environment.GetEnvString("COMPOSE_FILE", "local-test-data/docker-compose.yml"), []byte(NewContent), 0644)
	if err != nil {
		l.With("Error", err).Error("Error writing docker-compose.yml")
		return
	}

	l.Info("Compose File Updated")
}
