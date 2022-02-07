package no0045

import "fmt"

func jump(nums []int) int {
	//貪婪演算法，只取最好的那個
	if len(nums) == 1 {
		return 0
	}
	needChoose, canReach, step := 0, 0, 0
	for i, x := range nums {
		fmt.Println(i, x, needChoose, canReach, step)
		// i+x = key+value, 如果當下的步數大於原本紀錄的reach, 就換成i+x
		// 如果直接碰到nums的終點了，就直接回傳了，因為超出的reach point沒有意義
		if i+x > canReach {
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		// 替換下一個index前，讓canReach變成最新判斷的依據
		// step = index
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}
