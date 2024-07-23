package matchers

import (
	"errors"
	"feed-processor/search"
	"math/rand"
	"time"
)

type rssMatcher struct{}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func (m rssMatcher) Match(feed *search.FeedMetadata, searchTerm string) ([]*search.Result, error) {
	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	return nil, errors.New("not implemented")
}
