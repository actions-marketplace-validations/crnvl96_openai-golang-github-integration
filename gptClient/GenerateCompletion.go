package gptClient

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func GenerateCompletion(openai_api_key string, prompt string) gogpt.CompletionResponse {
	client := gogpt.NewClient(openai_api_key)
	context := context.Background()

	request := gogpt.CompletionRequest{
		Model:  "text-davinci-002",
		Prompt: "explain and give suggestions about this code:\n" + prompt,
		MaxTokens: 2048,
		Temperature: 0.5,
	}

	response, _ := client.CreateCompletion(context, request)

	return response
}