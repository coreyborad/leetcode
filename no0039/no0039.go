package no0039

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	thisLayer, result := []int{}, [][]int{}
	sort.Ints(candidates)
	dfs(candidates, target, 0, thisLayer, &result)
	return result
}

func dfs(cands []int, target int, startIndex int, thisLayer []int, result *[][]int) {
	if target <= 0 {
		// 因為前面有target-cands[i], 等於0代表是match的
		if target == 0 {
			ans := make([]int, len(thisLayer))
			copy(ans, thisLayer)
			*result = append(*result, ans)
		}
		return
	}
	for i := startIndex; i < len(cands); i++ {
		// 檢查數字本身就超過target就可以跳過了
		if cands[i] > target {
			break
		}
		thisLayer = append(thisLayer, cands[i])
		// 可以重複元素,i不需要做處理
		dfs(cands, target-cands[i], i, thisLayer, result)
		// 做完這個index後，剔除他，繼續往下窮舉
		thisLayer = thisLayer[:len(thisLayer)-1]
	}
}
