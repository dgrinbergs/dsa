package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BinarySearch(t *testing.T) {
	assert.Equal(t, 3, BinarySearch([]int{-1, 0, 2, 4, 6, 8}, 4))
}

func Test_BinarySearch2(t *testing.T) {
	assert.Equal(t, -1, BinarySearch([]int{-1, 0, 2, 4, 6, 8}, 3))
}

func Test_BinarySearch3(t *testing.T) {
	assert.Equal(t, 4, BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
}
