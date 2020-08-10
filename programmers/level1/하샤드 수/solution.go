package solution

func solution(x int) bool {
	tmp := x
	harshad := 0
	for tmp != 0 {
		harshad += tmp % 10
		tmp /= 10
	}
	if x%harshad == 0 {
		return true
	}
	return false
}
