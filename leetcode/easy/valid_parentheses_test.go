package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validParentheses(t *testing.T) {
	assert.True(t, validParentheses("[]"))
}

func Test_validParentheses2(t *testing.T) {
	assert.True(t, validParentheses("([{}])"))
}

func Test_validParentheses3(t *testing.T) {
	assert.False(t, validParentheses("[(])"))
}

func Test_validParentheses4(t *testing.T) {
	assert.False(t, validParentheses("("))
}

func Test_validParentheses5(t *testing.T) {
	assert.False(t, validParentheses("(("))
}

func Test_validParentheses6(t *testing.T) {
	assert.False(t, validParentheses("){"))
}
