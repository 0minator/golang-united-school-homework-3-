package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Sort(sort.IntSlice(keys))
	result = make([]string, 0, len(input))
	for _, key := range keys {
		result = append(result, input[key])
	}
	return result
}
