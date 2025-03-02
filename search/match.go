package search

import (
	"fmt"
	"log"
)

type Result struct {
	Title       string
	Description string
}

type Matcher interface {
	Match(feed *FeedMetadata, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *FeedMetadata, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Match(feed, searchTerm)
	if err != nil {
		log.Printf("Unable to match feed %s | err - %v", feed.Name, err)
		return
	}
	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("\n%s:\n%s\n", result.Title, result.Description)
	}
}
