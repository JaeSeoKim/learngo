package solution

import "sort"

func solution(n int64) int64 {
	var arr []int
	result := int64(0)
	for n != 0 {
		arr = append(arr, int(n%10))
		n /= 10
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] < arr[j] {
			return false
		}
		return true
	})
	for _, i := range arr {
		result *= 10
		result += int64(i)
	}
	return result
}
