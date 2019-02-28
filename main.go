package main

import (
	"log"
	"os"

	_ "github.com/goRss/matchers"
	"github.com/goRss/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	searchTerm := os.Args[1]

	search.Run(searchTerm)
}
