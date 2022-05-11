package homework

import "sort"

type Int64Slice []int64

func (x Int64Slice) Len() int           { return len(x) }
func (x Int64Slice) Less(i, j int) bool { return i < j }
func (x Int64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func reverse(input []int64) (result []int64) {
	slice2 := make([]int64, len(input))
	copy(slice2, input)
	sort.Sort(sort.Reverse(Int64Slice(slice2)))
	return slice2
}
