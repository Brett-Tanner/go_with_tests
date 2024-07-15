package main

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	sumTransactions := func(balance float64, t Transaction) float64 {
		if name == t.From {
			balance -= t.Sum
		} else if name == t.To {
			balance += t.Sum
		}

		return balance
	}

	return Reduce(transactions, sumTransactions, 0.0)
}

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

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initial B) B {
	result := initial
	for _, item := range collection {
		result = accumulator(result, item)
	}

	return result
}
