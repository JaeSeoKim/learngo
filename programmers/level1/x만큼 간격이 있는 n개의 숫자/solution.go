package solution

func solution(x int, n int) []int64 {
	var result []int64
	for i := 0; i < n; i++ {
		result = append(result, int64(x+x*i))
	}
	return result
}
