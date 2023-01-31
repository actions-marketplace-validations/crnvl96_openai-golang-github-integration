package githubClient

import (
	"context"

	"github.com/google/go-github/v32/github"
)

func CreateIssue(client *github.Client, context context.Context, title string, body string) {
	owner, repoName := getRepositoryNameAndOwner()

	issue := &github.IssueRequest{
		Title: github.String(title),
		Body:  github.String(body),
	}

	client.Issues.Create(context, owner, repoName, issue)
}