package medium

import (
	"slices"
)

// https://leetcode.com/problems/group-anagrams/description/

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		sorted := []rune(str)
		slices.Sort(sorted)

		if _, ok := groups[string(sorted)]; ok {
			groups[string(sorted)] = append(groups[string(sorted)], str)
			continue
		}

		groups[string(sorted)] = []string{str}
	}

	result := make([][]string, 0)

	for _, v := range groups {
		result = append(result, v)
	}

	return result
}
