package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/feed-metadata.json"

type FeedMetadata struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func LoadFeedMetadata() ([]*FeedMetadata, error) {
	file, err := os.Open(dataFile)
	defer file.Close() //this runs after executing this function, even on error/panic
	if err != nil {
		return nil, err
	}
	var feeds []*FeedMetadata
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
