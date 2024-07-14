package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("can sum numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9}, []int{1, 2, 3})
		want := []int{2, 9, 5}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplying all elements", func(t *testing.T) {
		multiply := func(a, b int) int {
			return a * b
		}

		got := Reduce([]int{2, 3, 4}, multiply, 1)
		want := 24

		if got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	})

	t.Run("concatenating strings", func(t *testing.T) {
		concatenate := func(a, b string) string {
			return a + b
		}

		got := Reduce([]string{"hello", " world"}, concatenate, "")
		want := "hello world"

		if got != want {
			t.Fatalf("got '%s' want '%s'", got, want)
		}
	})
}

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{From: "Brett", To: "Olga", Sum: 150.0},
		{From: "Vika", To: "Brett", Sum: 100.0},
	}

	AssertEqual(t, BalanceFor(transactions, "Brett"), -50)
	AssertEqual(t, BalanceFor(transactions, "Vika"), -100)
	AssertEqual(t, BalanceFor(transactions, "Olga"), 150)
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}
