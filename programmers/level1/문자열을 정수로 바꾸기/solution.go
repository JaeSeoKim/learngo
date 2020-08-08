package solution

func solution(s string) int {
	sign := 1
	integer := 0
	i := 0
	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			sign *= -1
		}
		i++
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		integer *= 10
		integer += int(s[i] - '0')
		i++
	}
	return integer * sign
}
