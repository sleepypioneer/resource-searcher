package analyse

import (
	"strings"
)

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
		words := strings.Split(p, " ")
		wordCount += len(words)
	}
	// return total word count for the article
	return wordCount
}

// FindTopics retuns a list of topics which the article includes
func FindTopics(content []string) (topics []string) {
	for _, p := range content {
		for _, keyword := range keywords {
			if strings.Contains(p, keyword) && !stringInSlice(keyword, topics) {
				topics = append(topics, keyword)
			}
		}
	}
	return topics
}
