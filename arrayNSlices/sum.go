package sum

// Blank identifier _

// Sum returns a sum of all integers of given array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}