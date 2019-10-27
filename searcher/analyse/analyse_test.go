package analyse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
