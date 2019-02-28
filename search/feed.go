package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed is a struct feed of notice in file data.json
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds unmarshal and return feeds
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
