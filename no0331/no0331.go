package no0331

import (
	"fmt"
	"strings"
)

// 以#結束
// 數字的個數比#的個數少一
// 左右兩個子節點都要以#結束
// 如果在還沒遍歷完就回傳validate = true，也算是false
func isValidSerialization(preorder string) bool {
	fmt.Println("====================")
	list := strings.Split(preorder, ",")
	i := 0
	result := validate(list, &i, "root") && i == len(list)
	fmt.Println(i, len(list))
	return result
}

func validate(preorder []string, i *int, direct string) bool {

	if len(preorder) == *i {
		return false
	}
	fmt.Println(preorder[*i], direct)
	if preorder[*i] == "#" {
		*i += 1
		return true
	}

	*i += 1
	left := validate(preorder, i, "left")
	right := validate(preorder, i, "right")
	fmt.Println(*i, left, right)
	return left && right
}
