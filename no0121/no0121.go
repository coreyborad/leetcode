package no0121

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	cur := prices[0]
	diff := 0
	for i := 1; i < len(prices); i++ {
		if cur > prices[i] {
			cur = prices[i]
		}
		if cur < prices[i] {
			temp := prices[i] - cur
			if temp > diff {
				diff = temp
			}
		}
	}
	return diff
}
