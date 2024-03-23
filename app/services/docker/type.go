package docker

import "strings"

type DCContainer struct {
	ID    string
	Name  string
	Image string
	Repo  string
	Tag   string
}

func getEnvUser(c DCContainer) string {

	EnvName := strings.ToUpper(c.Repo)
	EnvName = strings.ReplaceAll(EnvName, ".", "")

	return EnvName
}
