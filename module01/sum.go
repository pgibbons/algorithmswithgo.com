package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	results := 0
	// loop and find the number
	for _, value := range numbers {
		results += value
	}
	// return our results
	return results
}
