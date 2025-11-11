package medium

// https://leetcode.com/problems/product-of-array-except-self/description/

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	l := 1
	for i := range nums {
		if i > 0 {
			l *= nums[i-1]
		}

		result[i] = l
	}

	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i < len(nums)-1 {
			r *= nums[i+1]
		}

		result[i] *= r
	}

	return result
}
