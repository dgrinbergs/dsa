package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("Was it a car or a cat I saw?"))
}

func Test_isPalindrome2(t *testing.T) {
	assert.False(t, isPalindrome("tab a cat"))
}
