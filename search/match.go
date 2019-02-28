package search

import (
	"fmt"
	"log"
)

// Result is a struct
type Result struct {
	Field   string
	Content string
}

// Matcher is a interface of implementation Seach method
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is a goroutine process search
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

// Display a function show data
func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
