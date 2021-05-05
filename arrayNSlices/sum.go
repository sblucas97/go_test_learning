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

// SumAll returns a sum of all integers of given arrays
func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }
	// above the implementation being carefull with slices capacity

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails returns a sum of all integers of given arrays less than the head of each one
func SumAllTails(numbersToSum ...[]int) []int {
	var tails[][] int
	for _, numbers := range numbersToSum {
		var newElement []int

		if len(numbers) < 1 {
			newElement = []int {0}
		} else {
			newElement = numbers[1:]
		}

		tails = append(tails, newElement)
	}
	return SumAll(tails...)
}