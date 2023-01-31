package githubClient

import (
	"context"

	"github.com/google/go-github/v32/github"
)

func GetFileContent(client *github.Client, context context.Context, pr_id string, commitSHA string, fileName string) (string, error) {
	owner, repoName := getRepositoryNameAndOwner()
	fileContent, _, _, err := client.Repositories.GetContents(context, owner, repoName, fileName, &github.RepositoryContentGetOptions{Ref: commitSHA})

	content, _ := fileContent.GetContent()

	return content, err
}