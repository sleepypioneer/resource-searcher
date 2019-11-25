package resource

import (
	"log"
	"strings"
)

var (
	// keywords    = map[string]struct{}{"python": struct{}{}, "webservice": struct{}{}, "server": struct{}{}, "Django": struct{}{}, "list": struct{}{}}
	keywords    = []string{"python", "webservice", "server", "Django", "list"}
	punctuation = []string{".", ",", "!", ":", ";", "'", "/", "-", "_", "*", "&", "{", "}", "[", "]"}
)

// Resource stores the document and meta information for scraped data
type Resource struct {
	Document *Document `json:"document"`
	Meta     *Meta     `json:"meta"`
}

// type raw struct {
// 	meta meta `json:"meta"`
// 	document []byte `json:"document"`
// }

// Meta Data for scraped data
type Meta struct {
	HTTPStatus int    `json:"http_status"`
	URL        string `json:"url"`
}

// Document stores scraped data for document
type Document struct {
	Title     string   `json:"title"`
	Content   []string `json:"content"`
	Words     map[string]int
	WordCount int
	Topics    []string
}

// Analyse runs the various analysing logic on the give document
func (doc *Document) Analyse() error {
	err := doc.ProcessWords()
	if err != nil {
		return err
	}
	doc.Topics = FindTopics(doc.Words)
	log.Printf("doc word count: %d", doc.WordCount)
	log.Printf("doc topics: %v", doc.Topics)
	return nil
}

// ProcessWords returns total number of words in the article
func (doc *Document) ProcessWords() error {
	// for each paragraph in the article we count how many words it contains.
	wordsComplete := make(map[string]int)

	for _, p := range doc.Content {
		// remove all punction in paragraph
		for _, punc := range punctuation {
			p = strings.ReplaceAll(p, punc, "")
		}
		// split paragraph into words list
		words := strings.Fields(p)
		doc.WordCount += len(words)
		// wordsComplete = append(wordsComplete, words...)
		for _, word := range words {
			wordsComplete[word]++
		}
	}

	doc.Words = wordsComplete
	return nil
}

// FindTopics retuns a list of topics which the article includes
func FindTopics(words map[string]int) (topics []string) {
	for _, keyword := range keywords {
		if _, found := words[keyword]; found {
			topics = append(topics, keyword)
		}
	}
	return topics
}
