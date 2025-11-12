package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrix(t *testing.T) {
	assert.True(t, searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 10))
}

func Test_searchMatrix2(t *testing.T) {
	assert.False(t, searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 15))
}

func Test_searchMatrix3(t *testing.T) {
	assert.True(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
}
