package hard

// https://leetcode.com/problems/trapping-rain-water/description/

func trap(heights []int) int {
	prefixes := make([]int, len(heights))
	suffixes := make([]int, len(heights))

	// these are updated and referenced as we go
	lMax := 0
	rMax := 0

	for i, height := range heights {
		// take the greater of what we've seen and what we're currently seeing
		lMax = max(lMax, height)
		prefixes[i] = lMax

		// same as above but from the back of the array
		j := len(heights) - 1 - i
		rMax = max(rMax, heights[j])
		suffixes[j] = rMax
	}

	result := 0

	for i, height := range heights {
		result += min(prefixes[i], suffixes[i]) - height
	}

	return result
}
