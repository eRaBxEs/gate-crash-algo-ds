package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {

	if numbers == nil || len(numbers) == 0 {
		return 0
	}

	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
