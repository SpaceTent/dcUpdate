package DockerAPI

import (
	"context"
	l "log/slog"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RestartContainer(localContainer DCContainer) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		l.With("Error", err).Error("Error setting up docker client")
		return
	}
	defer cli.Close()

	noWaitTimeout := 0 // to not wait for the container to exit gracefully
	if err := cli.ContainerRestart(ctx, localContainer.ID, container.StopOptions{Timeout: &noWaitTimeout}); err != nil {
		l.With("Error", err).Error("Error in docker ContainerRestart")
		return
	}

	l.Info("Container " + localContainer.Name + " Restarted")
}
