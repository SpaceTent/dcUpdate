package docker

import "C"
import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	l "log/slog"
	"os"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"

	"dcupdate/app/environment"
)

func GetNewImage(localContainer DCContainer, NewTag string) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		l.With("Error", err).Error("Error setting up docker client")
		return
	}
	defer cli.Close()

	// does this Container Repo need a login?  Chances are it does
	// this means we look in the ENV for the username and password based on the Repo Name.
	// for example the GitHub Repo ghrc.io will look for the ENV variables GHRCIO_USERNAME and GHRCIO_PASSWORD
	if environment.GetEnvString(getEnvUser(localContainer)+"_USERNAME", "") == "" {
		l.With("Repo", localContainer.Repo).Error("No ENV variables found for Repo Username (e.g. GHCRIO_USERNAME)")
	}

	if environment.GetEnvString(getEnvUser(localContainer)+"_PASSWORD", "") == "" {
		l.With("Repo", localContainer.Repo).Error("No ENV variables found for Repo Password (e.g. GHCRIO_PASSWORD)")
	}

	authConfig := registry.AuthConfig{
		Username: environment.GetEnvString(getEnvUser(localContainer)+"_USERNAME", ""),
		Password: environment.GetEnvString(getEnvUser(localContainer)+"_PASSWORD", ""),
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		if err != nil {
			l.With("Error", err).Error("Error Unable to Decode JSON Auth Config")
			return
		}
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	reader, err := cli.ImagePull(ctx, localContainer.Name+":"+NewTag, image.PullOptions{RegistryAuth: authStr})
	if err != nil {
		l.With("Error", err).Error("Error setting up docker client")
		return
	}
	io.Copy(os.Stdout, reader)

}
