package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pkg/errors"
)

type resource struct {
	Document *document `json:"document"`
	Meta     *meta     `json:"meta"`
}

// type raw struct {
// 	meta meta `json:"meta"`
// 	document []byte `json:"document"`
// }

type meta struct {
	HTTPStatus int    `json:"http_status"`
	URL        string `json:"url"`
}

type document struct {
	Title     string   `json:"title"`
	Content   []string `json:"content"`
	wordCount int
	topics    []string
}

var (
	keywords    = []string{"python", "webservice", "server", "Django", "list"}
	punctuation = []string{".", ",", "!", ":", ";", "'", "/", "-", "_", "*", "&", "{", "}", "[", "]"}
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (s *service) scrape(url string) (*resource, error) {
	// read file for the moment
	// in the future: call the python scraper
	data, err := ioutil.ReadFile("./resource.json")
	if err != nil {
		return nil, err
	}

	res := &resource{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, errors.Wrap(err, "unmarshalling JSON")
	}
	return res, nil
}

func wordCount(content []string) int {
	wordCount := 0
	// for each paragraph in the article we count how many words it contains.
	for _, p := range content {
		// remove all punction in paragraph
		for _, punc := range punctuation {
			p = strings.ReplaceAll(p, punc, "")
		}
		// split paragraph into words list
		words := strings.Split(p, " ")
		wordCount += len(words)
	}
	// return total word count for the article
	return wordCount
}

func findTopics(content []string) (topics []string) {
	for _, p := range content {
		for _, keyword := range keywords {
			if strings.Contains(p, keyword) && !stringInSlice(keyword, topics) {
				topics = append(topics, keyword)
			}
		}
	}
	return topics
}

func analyse(doc *document) {
	// do something with your document
	doc.wordCount = wordCount(doc.Content)
	doc.topics = findTopics(doc.Content)
}

type service struct {
	// db *sql.DB
}

func (s *service) store(res *resource) error {
	// store to DB
	return nil
}

func main() {

	s := new(service)

	urls := []string{
		"stackoverflow.com",
	}

	for _, url := range urls {
		resource, err := s.scrape(url)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		// do we need an error to be returned?
		analyse(resource.Document)
		log.Printf("doc word count: %d", resource.Document.wordCount)
		log.Printf("doc topics: %v", resource.Document.topics)

		if err := s.store(resource); err != nil {
			log.Printf("ERROR: %v", err)
		}
	}

	log.Print("Bye")
}
