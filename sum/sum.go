package main

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	var results []int
	for _, slice := range slicesToSum {
		results = append(results, Sum(slice))
	}
	return results
}

func SumAllTails(slicesToSum ...[]int) []int {
	var results []int
	for _, slice := range slicesToSum {
		if len(slice) == 0 {
			results = append(results, 0)
			continue
		}
		results = append(results, Sum(slice[1:]))
	}
	return results
}
