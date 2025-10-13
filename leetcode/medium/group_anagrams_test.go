package medium

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams_Simple(t *testing.T) {
	assert.ElementsMatch(t, groupAnagrams([]string{""}), [][]string{{""}})
	assert.ElementsMatch(t, groupAnagrams([]string{"a"}), [][]string{{"a"}})
}

func TestGroupAnagrams_Complex(t *testing.T) {
	grouped := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

	for _, group := range grouped {
		slices.Sort(group)
	}

	assert.Contains(t, grouped, []string{"bat"})
	assert.Contains(t, grouped, []string{"nat", "tan"})
	assert.Contains(t, grouped, []string{"ate", "eat", "tea"})
}
