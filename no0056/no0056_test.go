package leetcode

import (
	"fmt"
	"testing"
)

type question56 struct {
	para56
	ans56
}

// para 是参数
// one 代表第一个参数
type para56 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans56 struct {
	one [][]int
}

func Test_Problem56(t *testing.T) {
	qs := []question56{
		{
			para56{
				one: [][]int{},
			},
			ans56{
				one: [][]int{},
			},
		},
		{
			para56{
				one: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			ans56{
				one: [][]int{{1, 6}, {8, 10}, {15, 18}},
			},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 56------------------------\n")

	for _, q := range qs {
		_, p := q.ans56, q.para56
		fmt.Printf("【input】:%v       【output】:%v\n", p, merge56(p.one))
	}
	fmt.Printf("\n\n\n")
}
