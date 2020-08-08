package solution

func solution(n int64) int64 {
	i := int64(0)
	for ; i*i < n; i++ {
	}
	if i*i == n {
		return (i + 1) * (i + 1)
	}
	return -1
}
