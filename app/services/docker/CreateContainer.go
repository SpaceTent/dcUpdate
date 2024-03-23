package docker

import (
	"context"
	"fmt"
	l "log/slog"
	"math/rand"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func CreateContainer(localContainer DCContainer, NewTag string) error {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		l.With("Error", err).Error("Error setting up docker client")
		return err
	}
	defer cli.Close()

	// Delete the old container
	err = cli.ContainerRemove(ctx, localContainer.ID, container.RemoveOptions{Force: true})
	if err != nil {
		l.With("Error", err).Error("Error Deleting Container")
		return err
	}

	// Create the new container
	l.Info("Creating Container " + localContainer.Name)
	cc := container.Config{
		Image: localContainer.Name + ":" + NewTag,
	}
	ch := container.HostConfig{}
	cn := network.NetworkingConfig{}

	// TODO: Hard Coded
	p := v1.Platform{
		Architecture: "amd64",
		OS:           "linux",
	}

	// TODO: Not to Future Self: This creates a container with NONE of the configurations that the old container had.  This is a problem.
	// need to find a way of copying the old container's configuration to the new container, or reading the config from the Docker Compose file.
	createdContainer, err := cli.ContainerCreate(ctx, &cc, &ch, &cn, &p, fmt.Sprintf("container-%d", rand.Intn(1000)))
	if err != nil {
		l.With("Error", err).Error("Error creating container")
		return err
	}

	// Start it up - any config options?
	cli.ContainerStart(ctx, createdContainer.ID, container.StartOptions{})

	l.Info("Container " + localContainer.Name + " Created")
	return nil
}
