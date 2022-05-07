package arraysandslices

func Sum(items []int) int {
	var sum int
	for _, number := range items {
		sum += number
	}
	return sum
}

// Sum all the elements of the array.
func SumAll(numbersToSum ...[]int) []int {
	var sumAll []int
	for _, arr := range numbersToSum {
		sumAll = append(sumAll, Sum(arr))
	}
	return sumAll
}

// Sum the tails of the array. The sub array
// excluding the head is called the tail.
func SumAllTails(numbersToSum ...[]int) []int {
	var result []int
	for _, array := range numbersToSum {
		if len(array) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(array[1:]))
		}
	}
	return result
}
