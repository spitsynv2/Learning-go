package arrays

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	result := []int{}

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	result := []int{}

	for _, numbers := range numbersToSum {
		if len(numbers) < 1 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(numbers[1:]))
		}
	}
	return result
}
