package leetcode

func searchInsert(nums []int, target int) int {
	result := -1
	if len(nums) <= 0 {
		return 0
	}
	for index, num := range nums {
		if num >= target {
			result = index
			break
		}
	}
	if len(nums) > 0 && result == -1 {
		return len(nums)
	} else {
		return result
	}

}
