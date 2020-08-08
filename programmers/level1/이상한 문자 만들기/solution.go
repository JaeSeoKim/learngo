package solution

func solution(s string) string {
	str := []rune(s)
	count := 0
	for i, c := range str {
		println(i, c)
		if count%2 == 0 && c >= 'a' && c <= 'z' {
			str[i] -= 'a' - 'A'
		}
		if count%2 == 1 && c >= 'A' && c <= 'Z' {
			str[i] += 'a' - 'A'
		}
		if c == ' ' {
			count = 0
		} else {
			count++
		}
	}
	return string(str)
}
