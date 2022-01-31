package no0763

import "fmt"

// 這題的目的是想辦法切最多partition出來，paritition沒有限制
// 先掃描所有字母最後一次出現的位置
// 接著再跑一次字串迴圈
// 當前掃到的index剛好等於已掃過的字母中，最大值，就可以切割一個最小的partition

func partitionLabels(s string) []int {
	result := []int{}
	start := 0
	scanMax := 0
	maxPosition := map[string]int{}
	for i := 0; i < len(s); i++ {
		maxPosition[string(s[i])] = i
	}
	fmt.Println(maxPosition)

	for i := 0; i < len(s); i++ {
		currMax := maxPosition[string(s[i])]
		scanMax = goMax(currMax, scanMax)

		if i == scanMax {
			fmt.Println(i, currMax, scanMax, string(s[i]))
			result = append(result, i-start+1)
			start = i + 1
		}
	}

	return result
}

func goMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
