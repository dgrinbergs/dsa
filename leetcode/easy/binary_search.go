package easy

// https://neetcode.io/problems/binary-search

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			r = mid - 1
			continue
		}

		l = mid + 1
	}

	return -1
}
