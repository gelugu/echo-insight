package main

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
)

func getCompletion(messages []openai.ChatCompletionMessage) string {
	log.Println("Getting completion")

	client := openai.NewClient(config.Openai.ApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4Turbo0125,
			Messages: messages,
		},
	)

	log.Println(
		"ChatCompletion response usage:",
		"PromptTokens="+string(rune(resp.Usage.PromptTokens)),
		"CompletionTokens="+string(rune(resp.Usage.CompletionTokens)),
		"TotalTokens="+string(rune(resp.Usage.TotalTokens)),
	)

	if err != nil {
		log.Fatal("ChatCompletion error:", err)
	}

	return resp.Choices[0].Message.Content
}
