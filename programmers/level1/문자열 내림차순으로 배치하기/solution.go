package solution

import (
	"sort"
	"strings"
)

type sortStr []string

func (a sortStr) Len() int           { return len(a) }
func (a sortStr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortStr) Less(i, j int) bool { return a[i] < a[j] }
func solution(s string) string {
	str := strings.Split(s, "")
	sort.Sort(sort.Reverse(sortStr(str)))
	return strings.Join(str, "")
}
