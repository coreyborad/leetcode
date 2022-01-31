package no0020

func isValid(s string) bool {
	// (({(())}{}))

	stack := []string{}
	for _, v := range s {
		switch string(v) {
		case "(", "[", "{":
			stack = append(stack, string(v))
		case ")", "]", "}":
			if len(stack) <= 0 {
				return false
			}
			if string(v) == ")" && stack[len(stack)-1] != "(" {
				return false
			}
			if string(v) == "]" && stack[len(stack)-1] != "[" {
				return false
			}
			if string(v) == "}" && stack[len(stack)-1] != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
