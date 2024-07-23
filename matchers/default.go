package matchers

import (
	"errors"
	"feed-processor/search"
	"fmt"
)

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	search.Register("default", matcher)
}

func (m defaultMatcher) Match(feed *search.FeedMetadata, _ string) ([]*search.Result, error) {
	return nil, errors.New(fmt.Sprintf("implement proper matcher for type: %s", feed.Type))
}
