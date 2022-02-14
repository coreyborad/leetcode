package no0055

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	// [0]
	// [0,2,3]
	// [2,3,1,1,4]
	// [3,0,8,2,0,0,1]
	max := 0
	// 貪婪演算法，抓最大就可以了，提前碰到length就直接跳出了
	for i := 0; i < len(nums)-1; i++ {
		// index+最大可踩步數
		step := i + nums[i]
		if step > max {
			max = step
		}
		if max >= len(nums)-1 {
			return true
		}
		// max如果已經跟不上當前index，就代表根本無法在前進了，直接跳出就可以了
		if max <= i {
			return false
		}
	}
	return false
}
