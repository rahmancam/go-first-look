package utils

// Add adds together multiple numbers
func Add(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}
