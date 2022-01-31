package no0017

func letterCombinations(digits string) []string {
	var q []string
	if len(digits) == 0 {
		return q
	}
	q = append(q, "")
	for _, v := range digits {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if v == '2' {
				q = append(q, cur+"a")
				q = append(q, cur+"b")
				q = append(q, cur+"c")
			}
			if v == '3' {
				q = append(q, cur+"d")
				q = append(q, cur+"e")
				q = append(q, cur+"f")
			}
			if v == '4' {
				q = append(q, cur+"g")
				q = append(q, cur+"h")
				q = append(q, cur+"i")
			}
			if v == '5' {
				q = append(q, cur+"j")
				q = append(q, cur+"k")
				q = append(q, cur+"l")
			}
			if v == '6' {
				q = append(q, cur+"m")
				q = append(q, cur+"n")
				q = append(q, cur+"o")
			}
			if v == '7' {
				q = append(q, cur+"p")
				q = append(q, cur+"q")
				q = append(q, cur+"r")
				q = append(q, cur+"s")
			}
			if v == '8' {
				q = append(q, cur+"v")
				q = append(q, cur+"u")
				q = append(q, cur+"t")
			}
			if v == '9' {
				q = append(q, cur+"w")
				q = append(q, cur+"x")
				q = append(q, cur+"y")
				q = append(q, cur+"z")
			}
		}
	}
	return q
}
