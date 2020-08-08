package solution

func solution(a int, b int) int64 {
	var result int64 = 0
	if a > b {
		tmp := a
		a = b
		b = tmp
	}
	for i := a; i <= b; i++ {
		result += int64(i)
	}
	return result
}
