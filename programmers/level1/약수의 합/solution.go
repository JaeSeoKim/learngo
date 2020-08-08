package solution

func solution(n int) int {
	result := 0
	i := 1
	for ; i*i < n; i++ {
		if n%i == 0 {
			result += i
			result += n / i
		}
	}
	if i*i == n {
		result += i
	}
	return result
}
