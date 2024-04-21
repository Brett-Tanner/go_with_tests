package main

import (
	"testing"
)

const (
	word       = "test"
	definition = "this is a test"
)

func TestAdd(t *testing.T) {
	t.Run("adding a word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("adding a duplicate word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new definition")

		assertError(t, err, WordExistsErr)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)
		_, err := dictionary.Search(word)

		assertError(t, err, NotFoundErr)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, NotFoundErr)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{word: definition}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(word)
		want := definition

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

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, NotFoundErr)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should have found the word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
