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
	doc.WordCount = WordCount(doc.Content)
	doc.Topics = FindTopics(doc.Content)
	log.Printf("doc word count: %d", doc.WordCount)
	log.Printf("doc topics: %v", doc.Topics)
	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// WordCount returns total number of words in the article
func WordCount(content []string) int {
	wordCount := 0
	// for each paragraph in the article we count how many words it contains.
	for _, p := range content {
		// remove all punction in paragraph
		for _, punc := range punctuation {
			p = strings.ReplaceAll(p, punc, "")
		}
		// split paragraph into words list
		words := strings.Fields(p)
		wordCount += len(words)
	}
	// return total word count for the article
	return wordCount
}

// FindTopics retuns a list of topics which the article includes
func FindTopics(content []string) (topics []string) {
	var words []string
	for _, p := range content {
		words = append(words, strings.Fields(p)...)
	}

	mb := make(map[string]struct{}, len(words))
	for _, x := range words {
		mb[x] = struct{}{}
	}

	for _, keyword := range keywords {
		if _, found := mb[keyword]; found {
			topics = append(topics, keyword)
		}
	}
	return topics
}
