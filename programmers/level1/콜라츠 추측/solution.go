package solution

func solution(num int) int {
	cnt := 0
	for num != 1 {
		if cnt >= 500 {
			return -1
		}
		if num%2 == 0 {
			num /= 2
		} else {
			num *= 3
			num++
		}
		cnt++
	}
	return cnt
}
