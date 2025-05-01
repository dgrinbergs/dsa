package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	assert.False(t, containsDuplicate([]int{1, 2, 3}))
	assert.True(t, containsDuplicate([]int{1, 1, 3}))
}
