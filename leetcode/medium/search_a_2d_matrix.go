package medium

import (
	"github.com/dgrinbergs/dsa/leetcode/easy"
)

// https://neetcode.io/problems/search-2d-matrix

func searchMatrix(matrix [][]int, target int) bool {
	l := 0
	r := len(matrix) - 1

	for l <= r {
		mid := (l + r) / 2
		minimum := matrix[mid][0]
		maximum := matrix[mid][len(matrix[mid])-1]

		if target > minimum && target > maximum {
			l = mid + 1
			continue
		}

		if target < minimum && target < maximum {
			r = mid - 1
			continue
		}

		return easy.BinarySearch(matrix[mid], target) != -1
	}

	return false
}
