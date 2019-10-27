package analyse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result []string

func Test_WordCount(t *testing.T) {
	expectedWordCount := 9
	content := []string{"python is the best!", "Especially Django - I love it."}

	assert.Equal(t, expectedWordCount, WordCount(content))
}

func Test_FindTopics(t *testing.T) {
	expectedTopics := []string{"python", "Django"}
	content := []string{"python is the best!", "Especially Django - I love it."}

	assert.Equal(t, expectedTopics, FindTopics(content))
}

func Benchmark_FindTopics(b *testing.B) {
	var topics []string
	content := []string{"python is the best!", "Especially Django - I love it."}
	// run the FindTopics function b.N times
	for n := 0; n < b.N; n++ {
		topics = FindTopics(content)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
	result = topics
}
