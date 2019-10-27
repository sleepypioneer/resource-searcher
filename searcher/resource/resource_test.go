package resource

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var (
	// MockArticle is used for testing packages
	mockArticle = &Resource{}
)

func init() {
	data, err := ioutil.ReadFile("../mocks/mock_article.json")
	if err != nil {
		log.Info().Err(err).Msg("reading from mock article file")
	}

	if err := json.Unmarshal(data, mockArticle); err != nil {
		log.Info().Err(err).Msg("unmarshalling JSON")
	}
}

func Test_Analyse(t *testing.T) {
	expectedWordCount := 5838
	expectedTopics := []string{"python", "Django", "list"}
	err := mockArticle.Document.Analyse()

	assert.NoError(t, err)
	assert.Equal(t, expectedWordCount, mockArticle.Document.WordCount)
	assert.Equal(t, expectedTopics, mockArticle.Document.Topics)
}
