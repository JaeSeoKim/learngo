package solution

func solution(n int) string {
	watermelon := ""
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			watermelon += "수"
		} else {
			watermelon += "박"
		}
	}
	return watermelon
}
