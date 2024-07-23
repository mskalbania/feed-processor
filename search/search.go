package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {

	feeds, err := LoadFeedMetadata() //declaration + int operator instead of var
	if err != nil {
		log.Fatalf("Unable to retrieve feeds - %v", err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup //declaration and init to zero value
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *FeedMetadata) { //passed by value (feed as a mem address) so it is not shared with other coroutines
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done() //accessed via closure to share between coroutines
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait() //new go routine so Display can be called and results are printed as soon as arrive
		close(results)   //terminates Display for range loop and allows program to exit
	}()

	Display(results)
}

func Register(name string, matcher Matcher) {
	matchers[name] = matcher
	log.Printf("Registered matcher %s", name)
}
