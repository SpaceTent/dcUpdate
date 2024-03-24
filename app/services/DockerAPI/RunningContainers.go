package DockerAPI

import (
	"context"
	l "log/slog"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func GetRunningContainers() ([]DCContainer, error) {

	out := []DCContainer{}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		l.With("Error", err).Error("Error setting up docker client")
		return out, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		l.With("Error", err).Error("Error getting containers")
		return out, err
	}

	for _, c := range containers {
		c := DCContainer{
			ID:    c.ID,
			Image: c.Image,
		}
		nameSplit := strings.Split(c.Image, ":")
		if len(nameSplit) > 1 {
			c.Name = nameSplit[0]
			c.Tag = nameSplit[1]
		}
		nameSplit = strings.Split(c.Image, "/")
		if len(nameSplit) > 1 {
			c.Repo = nameSplit[0]
		}

		out = append(out, c)
		l.Info("Container: " + c.Name + " Tag: " + c.Tag + " Repo: " + c.Repo)
	}

	return out, nil
}
