package leetcode

// https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
	remainders := make(map[int]int)

	for i, n := range nums {
		for k, v := range remainders {
			if v == target-n && k != i {
				return []int{k, i}
			}
		}

		remainders[i] = n
	}

	panic("no solution found")
}
