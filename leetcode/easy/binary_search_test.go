package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binarySearch(t *testing.T) {
	assert.Equal(t, 3, binarySearch([]int{-1, 0, 2, 4, 6, 8}, 4))
}

func Test_binarySearch2(t *testing.T) {
	assert.Equal(t, -1, binarySearch([]int{-1, 0, 2, 4, 6, 8}, 3))
}

func Test_binarySearch3(t *testing.T) {
	assert.Equal(t, 4, binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
}
