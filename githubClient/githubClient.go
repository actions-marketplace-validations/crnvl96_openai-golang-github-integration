package githubClient

import (
	"context"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func GenerateClient(token string) (*github.Client, context.Context)  {
	context := context.Background()

	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tokenClient := oauth2.NewClient(context, tokenSource)
	client := github.NewClient(tokenClient)

	return client, context
}