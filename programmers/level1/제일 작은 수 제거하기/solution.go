package solution

func solution(arr []int) []int {
	var result []int
	min := arr[0]
	if len(arr) == 1 {
		return []int{-1}
	}
	for _, value := range arr {
		if min > value {
			min = value
		}
	}
	for _, value := range arr {
		if min != value {
			result = append(result, value)
		}
	}
	return result
}
