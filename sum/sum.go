package main

func SumAll(slicesToSum ...[]int) []int {
	sumArr := func(accumulator, x []int) []int {
		return sumFromOffset(0, accumulator, x)
	}

	return Reduce(slicesToSum, sumArr, []int{})
}

func SumAllTails(slicesToSum ...[]int) []int {
	sumTail := func(accumulator, x []int) []int {
		return sumFromOffset(1, accumulator, x)
	}

	return Reduce(slicesToSum, sumTail, []int{})
}

func sumFromOffset(offset int, accumulator, arr []int) []int {
	if len(arr) == 0 {
		return append(accumulator, 0)
	}

	return append(accumulator, Sum(arr[offset:]))
}

func Sum(numbers []int) int {
	return Reduce(numbers, add, 0)
}

func add(a, b int) int {
	return a + b
}

func Reduce[T any](collection []T, accumulator func(T, T) T, initial T) T {
	result := initial
	for _, item := range collection {
		result = accumulator(result, item)
	}

	return result
}
