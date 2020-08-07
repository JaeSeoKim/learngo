package solution

import "sort"

func solution(arr []int, divisor int) []int {
	var results []int
	for _, val := range arr {
		if val%divisor == 0 {
			results = append(results, val)
		}
	}
	sort.Ints(results)
	if results == nil {
		return append(results, -1)
	}
	return results
}
