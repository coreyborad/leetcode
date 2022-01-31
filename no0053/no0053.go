package leetcode

func maxSubArray(nums []int) int {
	// [-2,1,-3,4,-1,2,1,-5,4]
	// if last current sum > 0 keep add current index
	// else lest current sum equal current index
	// and you will get new current sum, to compare result and current sum

	cur := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if cur > 0 {
			cur += nums[i]
		} else {
			cur = nums[i]
		}

		if cur > result {
			result = cur
		}
	}
	return result
}
