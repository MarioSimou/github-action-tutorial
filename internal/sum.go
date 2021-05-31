package internal

// Sum accepts a list of integers and calculates the sum
func Sum(nums ...int) int {
	var sum = 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
