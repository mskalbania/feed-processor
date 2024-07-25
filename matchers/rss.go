package matchers

import (
	"encoding/xml"
	"errors"
	"feed-processor/search"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type rssMatcher struct{}

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func (m rssMatcher) Match(feed *search.FeedMetadata, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result
	log.Printf("Pooling rss [%s] feed from %s", feed.Name, feed.URI)
	document, err := fetchRssDocument(feed)
	if err != nil {
		return nil, err
	}
	for _, channelItem := range document.Channel.Item {
		matched, err := match(searchTerm, channelItem)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, toResult(channelItem))
		}
	}
	return results, nil
}

func fetchRssDocument(feed *search.FeedMetadata) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New(fmt.Sprintf("URI not provided in: %s", feed.Name))
	}
	rs, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error while calling %s, code: %d", feed.URI, rs.StatusCode))
	}
	var document rssDocument
	err = xml.NewDecoder(rs.Body).Decode(&document)
	return &document, err
}

func match(searchTerm string, channelItem item) (bool, error) {
	matchedInTitle, err := regexp.MatchString(searchTerm, channelItem.Title)
	matchedInDescription, err := regexp.MatchString(searchTerm, channelItem.Description)
	return matchedInTitle || matchedInDescription, err
}

func toResult(channelItem item) *search.Result {
	return &search.Result{
		Title:       channelItem.Title,
		Description: channelItem.Description,
	}
}
