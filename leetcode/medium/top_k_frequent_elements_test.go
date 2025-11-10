package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_topKFrequent(t *testing.T) {
	topK := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	assert.Contains(t, topK, 1)
	assert.Contains(t, topK, 2)
}

func Test_topKFrequent2(t *testing.T) {
	topK := topKFrequent([]int{1}, 1)
	assert.Contains(t, topK, 1)
}

func Test_topKFrequent3(t *testing.T) {
	topK := topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)
	assert.Contains(t, topK, 1)
	assert.Contains(t, topK, 2)
}
