package solution

import (
	"sort"
)

func solution(array []int, commands [][]int) []int {
	var result []int
	for _, command := range commands {
		tmp := make([]int, len(array))
		copy(tmp, array)
		result = append(result, findNumK(tmp, command))
	}
	return result
}

func findNumK(array []int, command []int) int {
	array = array[command[0]-1 : command[1]]
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	return array[command[2]-1]
}
