package DockerAPI

import (
	"context"
	l "log/slog"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func StopContainer(localContainer DCContainer, NewTag string) error {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		l.With("Error", err).Error("Error setting up docker client")
		return err
	}
	defer cli.Close()

	cli.ContainerStop(ctx, localContainer.ID, container.StopOptions{})

	return nil
}
