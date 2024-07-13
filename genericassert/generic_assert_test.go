package genericassert_test

import (
	"testing"

	"github.com/Brett-Tanner/go_with_tests/genericassert"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		genericassert.AssertEqual(t, 1, 1)
		genericassert.AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		genericassert.AssertEqual(t, "hello", "hello")
		genericassert.AssertNotEqual(t, "hello", "yoyo")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		intStack := new(genericassert.Stack[int])

		genericassert.AssertTrue(t, intStack.IsEmpty())

		intStack.Push(69)
		genericassert.AssertFalse(t, intStack.IsEmpty())

		intStack.Push(100)
		val, _ := intStack.Pop()
		genericassert.AssertEqual(t, val, 100)
		val, _ = intStack.Pop()
		genericassert.AssertEqual(t, val, 69)
		genericassert.AssertTrue(t, intStack.IsEmpty())

		intStack.Push(2)
		intStack.Push(3)
		firstNum, _ := intStack.Pop()
		secondNum, _ := intStack.Pop()
		genericassert.AssertEqual(t, 5, firstNum+secondNum)
	})

	t.Run("string stack", func(t *testing.T) {
		stringStack := new(genericassert.Stack[string])

		genericassert.AssertTrue(t, stringStack.IsEmpty())

		stringStack.Push("sixty nine")
		genericassert.AssertFalse(t, stringStack.IsEmpty())

		stringStack.Push("one hunnid")
		val, _ := stringStack.Pop()
		genericassert.AssertEqual(t, val, "one hunnid")
		val, _ = stringStack.Pop()
		genericassert.AssertEqual(t, val, "sixty nine")
		genericassert.AssertTrue(t, stringStack.IsEmpty())

		stringStack.Push("2")
		stringStack.Push("3")
		firstChar, _ := stringStack.Pop()
		secondChar, _ := stringStack.Pop()
		genericassert.AssertEqual(t, firstChar+secondChar, "32")
	})
}
