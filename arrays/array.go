package arrays

func Sum(val []int) int {
	var sum = 0
	//for i := 0; i < len(val); i++ {
	//	sum += val[i]
	//}
	for _, num := range val {
		sum += num
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
