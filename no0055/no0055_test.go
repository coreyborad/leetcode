package no0055

import (
	"fmt"
	"testing"
)

type question55 struct {
	para55
	ans55
}

// para 是参数
// one 代表第一个参数
type para55 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans55 struct {
	one bool
}

func Test_Problem55(t *testing.T) {

	qs := []question55{
		// {
		// 	para55{[]int{0}},
		// 	ans55{true},
		// },
		{
			para55{[]int{0, 2, 3}},
			ans55{false},
		},
		{
			para55{[]int{2, 3, 1, 1, 4}},
			ans55{true},
		},
		{
			para55{[]int{3, 2, 1, 0, 4}},
			ans55{false},
		},
		{
			para55{[]int{3, 0, 8, 2, 0, 0, 1}},
			ans55{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 55------------------------\n")

	for _, q := range qs {
		a, p := q.ans55, q.para55
		fmt.Printf("【input】:%v       【output】:%v     【answer】:%v\n", p, canJump(p.one), a)
	}
	fmt.Printf("\n\n\n")
}