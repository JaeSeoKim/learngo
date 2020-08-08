package solution

import "strconv"

func solution(seoul []string) string {
	for i, val := range seoul {
		if val == "Kim" {
			return "김서방은 " + strconv.Itoa(i) + "에 있다"
		}
	}
	return ""
}
