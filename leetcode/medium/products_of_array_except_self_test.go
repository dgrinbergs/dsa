package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf(t *testing.T) {
	result := productExceptSelf([]int{1, 2, 3, 4})
	assert.Equal(t, []int{24, 12, 8, 6}, result)
}

func Test_productExceptSelf2(t *testing.T) {
	result := productExceptSelf([]int{-1, 1, 0, -3, 3})
	assert.Equal(t, []int{0, 0, 9, 0, 0}, result)
}
