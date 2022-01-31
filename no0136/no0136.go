package no0136

func singleNumber(nums []int) int {
	mapSave := map[int]int{}
	for i := 0; i < len(nums); i++ {
		mapSave[nums[i]] += 1
	}
	for key, v := range mapSave {
		if v == 1 {
			return key
		}
	}
	return 0
}
