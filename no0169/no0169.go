package no0169

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	r := 0
	ti := -1
	for index, count := range m {
		if count > r {
			r = count
			ti = index
		}
	}
	return ti
}
