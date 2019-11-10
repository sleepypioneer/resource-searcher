package resource

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

const mockArticlePath = "./mock_article.json"

var (
	// MockArticle is used for testing packages
	mockArticle = &Resource{}
	result      []string
)

func init() {
	data, err := ioutil.ReadFile(mockArticlePath)
	if err != nil {
		log.Info().Err(err).Msg("reading from mock article file")
	}

	if err := json.Unmarshal(data, mockArticle); err != nil {
		log.Info().Err(err).Msg("unmarshalling JSON")
	}
}

func Test_Analyse(t *testing.T) {
	expectedWordCount := 5841
	expectedTopics := []string{"python", "Django", "list"}
	err := mockArticle.Document.Analyse()

	assert.NoError(t, err)
	assert.Equal(t, expectedWordCount, mockArticle.Document.WordCount)
	assert.Equal(t, expectedTopics, mockArticle.Document.Topics)
}

func Test_WordCount(t *testing.T) {
	expectedWordCount := 5841
	content := mockArticle.Document.Content

	assert.Equal(t, expectedWordCount, WordCount(content))
}

func Test_FindTopics(t *testing.T) {
	expectedTopics := []string{"python", "Django", "list"}
	content := mockArticle.Document.Content
	assert.Equal(t, expectedTopics, FindTopics(content))
}

func Benchmark_FindTopics(b *testing.B) {
	var topics []string
	content := mockArticle.Document.Content
	// run the FindTopics function b.N times
	for n := 0; n < b.N; n++ {
		topics = FindTopics(content)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
	result = topics
}
