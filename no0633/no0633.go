package no0633

import "math"

func judgeSquareSum(c int) bool {
	for i, j := 0, int(math.Sqrt(float64(c))); i <= j; {
		if sum := i*i + j*j; sum < c {
			i++
		} else if sum > c {
			j--
		} else {
			return true
		}
	}
	return false
}
