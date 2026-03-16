package golang

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	start := 0
	sign := 1

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for index := start; index < len(s); index++ {
		if s[index] < '0' || s[index] > '9' {
			return 0
		}
		result = result*10 + int(s[index]-'0')
	}
	return result * sign
}
