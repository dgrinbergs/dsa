package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trap(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func Test_trap2(t *testing.T) {
	assert.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5}))
}

func Test_trap3(t *testing.T) {
	assert.Equal(t, 4, trap([]int{4, 0, 4}))
}
