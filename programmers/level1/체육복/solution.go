package solution

func solution(n int, lost []int, reserve []int) int {
	clothes := make([]int, n)
	for i := 0; i < n; i++ {
		clothes[i] = 1
	}
	for _, i := range lost {
		clothes[i-1]--
	}
	for _, i := range reserve {
		clothes[i-1]++
	}
	for i := 0; i < n; i++ {
		if clothes[i] == 2 {
			if i != 0 && clothes[i-1] == 0 {
				clothes[i] = 1
				clothes[i-1] = 1
			} else if i != n-1 && clothes[i+1] == 0 {
				clothes[i] = 1
				clothes[i+1] = 1
			}
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		if clothes[i] != 0 {
			count++
		}
	}
	return (count)
}
