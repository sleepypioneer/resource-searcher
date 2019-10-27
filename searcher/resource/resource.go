package resource

import (
	"log"

	"github.com/sleepypioneer/resource-searcher/searcher/analyse"
)

type Resource struct {
	Document *Document `json:"document"`
	Meta     *Meta     `json:"meta"`
}

// type raw struct {
// 	meta meta `json:"meta"`
// 	document []byte `json:"document"`
// }

type Meta struct {
	HTTPStatus int    `json:"http_status"`
	URL        string `json:"url"`
}

type Document struct {
	Title     string   `json:"title"`
	Content   []string `json:"content"`
	WordCount int
	Topics    []string
}

// Analyse runs the various analysing logic on the give document
func (doc *Document) Analyse() error {
	// do something with your document
	doc.WordCount = analyse.WordCount(doc.Content)
	doc.Topics = analyse.FindTopics(doc.Content)
	log.Printf("doc word count: %d", doc.WordCount)
	log.Printf("doc topics: %v", doc.Topics)
	return nil
}
