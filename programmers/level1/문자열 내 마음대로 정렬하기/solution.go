package solution

import (
	"sort"
	str "strings"
)

func solution(strings []string, n int) []string {
	sort.Slice(strings, func(i, j int) bool {
		if strings[i][n] > strings[j][n] {
			return false
		}
		if strings[i][n] == strings[j][n] {
			return (str.Compare(strings[i], strings[j]) != 1)
		}
		return true
	})
	return strings
}
