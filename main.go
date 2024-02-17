package main

import (
	"github.com/sashabaranov/go-openai"
	"log"
)

func main() {
	log.Println("Starting")

	var completionMessages = []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: config.Openai.MainPrompt,
		},
	}
	if config.Openai.CustomSyntax != "" {
		completionMessages = append(completionMessages, openai.ChatCompletionMessage{
			Role:    "system",
			Content: config.Openai.CustomSyntax,
		})
	}

	log.Println("Parsing projects")
	for _, projectId := range config.Gitlab.ProjectIds {
		project := getProject(projectId)
		projectInfo := formatRepository(project)

		commits := getCommits(projectId)
		commitsPrompt := ""
		for _, commit := range commits {
			commitsPrompt += formatCommit(*commit) + "\n"
		}

		completionMessages = append(completionMessages, openai.ChatCompletionMessage{
			Role:    "user",
			Content: "Repository: " + projectInfo + "\nCommits:\n" + commitsPrompt,
		})
		completion := getCompletion(completionMessages)

		// TODO: make sure completion is in utf-8
		if len(completion) > 4096 {
			log.Println("Completion is too long, sending in parts")
			parts := splitByLength(completion, 4096)
			for _, part := range parts {
				sendMessage(config.Telegram.Token, config.Telegram.ChatId, part)
			}
		} else {
			sendMessage(config.Telegram.Token, config.Telegram.ChatId, completion)
		}
	}

	log.Println("Done")
}

func splitByLength(s string, length int) []string {
	var parts []string
	for len(s) > length {
		i := length
		for s[i] != '\n' {
			i--
		}
		parts = append(parts, s[:i])
		s = s[i+1:]
	}
	parts = append(parts, s)
	return parts
}
