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
	expectedWordCount := 2949
	expectedTopics := []string{"python", "django", "list"}
	err := mockArticle.Document.Analyse()

	assert.NoError(t, err)
	assert.Equal(t, expectedWordCount, mockArticle.Document.WordCount)
	assert.Equal(t, expectedTopics, mockArticle.Document.Topics)
}

func Test_FindTopics(t *testing.T) {
	expectedTopics := []string{"python", "django", "list"}
	words := mockArticle.Document.Words

	assert.Equal(t, expectedTopics, FindTopics(words))
}

func Benchmark_FindTopics(b *testing.B) {
	var topics []string
	words := mockArticle.Document.Words
	// run the FindTopics function b.N times
	for n := 0; n < b.N; n++ {
		topics = FindTopics(words)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
	result = topics
}
