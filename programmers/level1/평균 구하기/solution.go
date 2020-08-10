package solution

func solution(arr []int) float64 {
	result := float64(0)
	for _, val := range arr {
		result += float64(val)
	}
	return result / float64(len(arr))
}
