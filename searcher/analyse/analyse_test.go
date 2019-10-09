package analyse

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/sleepypioneer/resource-searcher/searcher/resource"
	"github.com/stretchr/testify/assert"
)

var (
	mockArticle = &resource.Resource{}
)

func init() {
	data, err := ioutil.ReadFile("./mock_article.json")
	if err != nil {
		log.Info().Err(err).Msg("reading from mock article file")
	}

	if err := json.Unmarshal(data, mockArticle); err != nil {
		log.Info().Err(err).Msg("unmarshalling JSON")
	}
}

func Test_WordCount(t *testing.T) {
	expectedWordCount := 10
	content := []string{"python is the best!", "Especially Django - I love it."}

	assert.Equal(t, expectedWordCount, WordCount(content))
}

func Test_FindTopics(t *testing.T) {
	expectedTopics := []string{"python", "Django"}
	content := []string{"python is the best!", "Especially Django - I love it."}

	assert.Equal(t, expectedTopics, FindTopics(content))
}

func Test_Analyse(t *testing.T) {
	expectedWordCount := 5833
	expectedTopics := []string{"python", "list", "Django"}

	Analyse(mockArticle.Document)

	assert.Equal(t, expectedWordCount, mockArticle.Document.WordCount)
	assert.Equal(t, expectedTopics, mockArticle.Document.Topics)
}
