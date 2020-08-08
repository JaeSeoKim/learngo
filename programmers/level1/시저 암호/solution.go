package solution

func solution(s string, n int) string {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			str[i] = rune((s[i]+byte(n)-'a')%26 + 'a')
		}
		if s[i] >= 'A' && s[i] <= 'Z' {

			str[i] = rune((s[i]+byte(n)-'a')%26 + 'a')
		}
	}
	return string(str)
}
