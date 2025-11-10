package medium

// https://leetcode.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)

	for _, num := range nums {
		seen[num]++
	}

	seenFreq := make([][]int, len(nums))
	for num, freq := range seen {
		seenFreq[freq-1] = append(seenFreq[freq-1], num)
	}

	result := make([]int, k)
	index := 0
	for i := len(seenFreq) - 1; i >= 0; i-- {
		if len(seenFreq[i]) < 1 {
			continue
		}

		for _, num := range seenFreq[i] {
			if index > k-1 {
				return result
			}

			result[index] = num
			index++
		}
	}

	return result
}
