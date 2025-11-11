package easy

import (
	"strings"
)

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	s = strings.ToLower(s)

	for l < r {
		left := s[l]
		right := s[r]

		if (left < 'a' || left > 'z') && (left < '0' || left > '9') {
			l++
			continue
		}

		if (right < 'a' || right > 'z') && (right < '0' || right > '9') {
			r--
			continue
		}

		if left != right {
			return false
		}

		l++
		r--
	}

	return true
}
