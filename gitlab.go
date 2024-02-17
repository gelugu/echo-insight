package main

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"log"
	"time"
)

// TODO: branches, tags, and merge requests

func getCommits(projectId int) []*gitlab.Commit {
	log.Println("Getting commits for project", projectId, "since", config.Days, "days")

	git, err := gitlab.NewClient(config.Gitlab.ApiKey)
	if err != nil {
		log.Fatal("Error creating client:", err)
	}

	currentTime := time.Now()
	oneDayAgo := currentTime.AddDate(0, 0, config.Days*-1)

	commits, _, err := git.Commits.ListCommits(projectId, &gitlab.ListCommitsOptions{
		Since: &oneDayAgo,
		Until: &currentTime,
	})
	if err != nil {
		log.Fatal("Error getting commits:", err)
	}

	log.Println("Got", len(commits), "commits")

	return commits
}

func formatCommit(commit gitlab.Commit) string {
	return fmt.Sprintf(
		"---\n"+
			"Autor: %s\n"+
			"Title: %s\n"+
			"Date: %s\n"+
			"URL: %s\n"+
			commit.AuthorName,
		commit.Title,
		commit.CommittedDate,
		commit.WebURL,
	)
}

func getProject(projectId int) *gitlab.Project {
	log.Println("Getting project", projectId)
	git, err := gitlab.NewClient(config.Gitlab.ApiKey)
	if err != nil {
		log.Fatal("Error creating client:", err)
	}

	repo, _, err := git.Projects.GetProject(projectId, nil)
	if err != nil {
		log.Fatal("Error getting project:", err)
	}

	return repo
}

func formatRepository(repo *gitlab.Project) string {
	return fmt.Sprintf(
		"%s: %s",
		repo.Name,
		repo.WebURL,
	)
}
