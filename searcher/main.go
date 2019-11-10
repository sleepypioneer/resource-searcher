package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"

	"github.com/sleepypioneer/resource-searcher/searcher/resource"
)

type service struct {
	// db *sql.DB
}

func (s *service) scrape(url string) (*resource.Resource, error) {
	// read file for the moment
	// in the future: call the python scraper
	data, err := ioutil.ReadFile("./resource.json")
	if err != nil {
		return nil, err
	}

	res := &resource.Resource{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, errors.Wrap(err, "unmarshalling JSON")
	}
	return res, nil
}

func (s *service) store(res *resource.Resource) error {
	// store to DB
	return nil
}

func main() {

	s := new(service)

	urls := []string{
		"stackoverflow.com",
	}

	for _, url := range urls {
		scrapedResource, err := s.scrape(url)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		// do we need an error to be returned?
		err = scrapedResource.Document.Analyse()
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		if err := s.store(scrapedResource); err != nil {
			log.Printf("ERROR: %v", err)
		}
	}

	log.Print("Bye")
}
