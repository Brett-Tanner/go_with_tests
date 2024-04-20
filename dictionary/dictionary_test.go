package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	word := "test"
	value := "this is a test"

	dictionary := Dictionary{}
	dictionary.Add(word, value)

	want := value
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should have found the word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSearch(t *testing.T) {
	word := "test"
	value := "this is a test"
	dictionary := Dictionary{word: value}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(word)
		want := value

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := NotFoundErr.Error()

		if err == nil {
			t.Fatal("expected error but none occured")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
