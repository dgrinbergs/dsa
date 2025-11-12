package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	assert.Equal(t, maxArea([]int{1, 7, 2, 5, 4, 7, 3, 6}), 36)
}

func Test_maxArea2(t *testing.T) {
	assert.Equal(t, maxArea([]int{2, 2, 2}), 4)
}
