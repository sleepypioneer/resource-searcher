package analyse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result []string

func Test_FindTopics(t *testing.T) {
	expectedTopics := []string{"python", "Django"}
	words := map[string]int{"python": 1, "is": 1, "the": 1, "best": 1, "Especially": 1, "Django": 1, "I": 1, "love": 1, "it": 1}

	assert.Equal(t, expectedTopics, FindTopics(words))
}

func Benchmark_FindTopics(b *testing.B) {
	var topics []string
	words := map[string]int{"python": 1, "is": 1, "the": 1, "best": 1, "Especially": 1, "Django": 1, "I": 1, "love": 1, "it": 1}
	// run the FindTopics function b.N times
	for n := 0; n < b.N; n++ {
		topics = FindTopics(words)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
	result = topics
}
