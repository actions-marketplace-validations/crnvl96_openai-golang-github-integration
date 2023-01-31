package main

import (
	"os"

	"github.com/crnvl96/openai-code-review/githubClient"
	"github.com/crnvl96/openai-code-review/gptClient"
)


func main() {
	openai_api_key := os.Getenv("OPENAI_API_KEY")
	github_token := os.Getenv("GITHUB_TOKEN")
	github_pr_id := os.Getenv("GITHUB_PR_ID")

	client, context := githubClient.GenerateClient(github_token)

	commits, _ := githubClient.GetCommits(client, context, github_pr_id)

	for _, commit := range commits {
		files := commit.Files
		for _, file := range files {
			fileName := file.GetFilename()
			commitSHA := commit.GetSHA()

			content, _ := githubClient.GetFileContent(client, context, github_pr_id, commitSHA, fileName)

			codeReview := gptClient.GenerateCompletion(openai_api_key, content)

			codeReviewText := codeReview.Choices[0].Text

			githubClient.CreateIssue(client, context, fileName, codeReviewText)			
		}
	}
	
}
