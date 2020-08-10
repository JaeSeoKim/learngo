package solution

func solution(arr1 [][]int, arr2 [][]int) [][]int {
	var result [][]int

	for i := 0; i < len(arr1); i++ {
		var tmp []int
		for j := 0; j < len(arr1[i]); j++ {
			sum := arr1[i][j] + arr2[i][j]
			tmp = append(tmp, sum)
		}
		result = append(result, tmp)
	}
	return result
}
