package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	// lenNums := len(numbers)
	totalsums := make([]int, len(numbers))

	for i, nums := range numbers {
		totalsums[i] = Sum(nums)
	}

	return totalsums
}

func SumAllTails(numbers ...[]int) []int {
	var totalTails []int 

	for _, nums := range numbers {
		if len(nums) == 0 {
			totalTails = append(totalTails, 0)
		} else {
			tail := nums[1:]
			totalTails = append(totalTails, Sum(tail))
		}
	}

	return totalTails
}