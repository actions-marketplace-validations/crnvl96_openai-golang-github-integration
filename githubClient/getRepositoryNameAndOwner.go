package githubClient

import (
	"os"
	"strings"
)

func getRepositoryNameAndOwner() (owner string, repoName string) {
	githubRepository := os.Getenv("GITHUB_REPOSITORY")
	parts := strings.Split(githubRepository, "/")
	owner = parts[0]
	repoName = parts[1]
	return
}