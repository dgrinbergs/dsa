package medium

// https://neetcode.io/problems/max-water-container?list=neetcode150

func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	result := 0

	for l < r {
		area := (r - l) * min(heights[l], heights[r])
		result = max(area, result)

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return result
}
