package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

type resource struct {
	document *document
	meta     *meta
}

// type raw struct {
// 	meta meta `json:"meta"`
// 	document []byte `json:"document"`
// }

type meta struct {
	httpStatus int    `json:"http_status"`
	url        string `json : "url"`
}

type document struct {
	title     string `json: "title"`
	content   string `json : "content"`
	wordCount int
	topics    []string
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

func analyse(doc *document) {
	// do something with your document
	doc.wordCount = 42
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
		analyse(resource.document)

		if err := s.store(resource); err != nil {
			log.Printf("ERROR: %v", err)
		}
	}

	log.Print("Bye")
}
