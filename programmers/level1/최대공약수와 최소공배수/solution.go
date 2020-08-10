package solution

func solution(n int, m int) []int {
	return []int{gcd(n, m), n * m / gcd(n, m)}
}

func gcd(n, m int) int {
	for m != 0 {
		tmp := n % m
		n = m
		m = tmp
	}
	return n
}
