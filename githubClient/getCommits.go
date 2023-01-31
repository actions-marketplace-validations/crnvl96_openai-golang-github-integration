package githubClient

import (
	"context"
	"strconv"

	"github.com/google/go-github/v32/github"
)

func GetCommits(client *github.Client, context context.Context, pr_id string) ([]*github.RepositoryCommit, error) {
	pullRequestId, _ := strconv.Atoi(pr_id)
	owner, repoName := getRepositoryNameAndOwner()

	commits, _, err := client.PullRequests.ListCommits(context, owner, repoName, pullRequestId, &github.ListOptions{})

	return commits, err
}