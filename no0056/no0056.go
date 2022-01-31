package no0056

import (
	"fmt"
	"sort"
)

func merge56(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	result := [][]int{
		intervals[0],
	}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		if last[1] >= intervals[i][0] {
			fmt.Println(last[1], intervals[i][1])
			result[len(result)-1][1] = max(last[1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}
