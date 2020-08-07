package solution

func solution(s string) string {
	if len(s)%2 == 1 {
		return s[len(s)/2 : (len(s)/2)+1]
	}
	return s[(len(s)/2)-1 : (len(s)/2)+1]
}
