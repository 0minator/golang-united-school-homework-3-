package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	for i := 0; i < len(input); i++ {
		result[len(input)-i-1] = input[i]
	}
	return result
}
