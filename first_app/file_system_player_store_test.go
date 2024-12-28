package main

import (
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	jsonLeague := `[
		{"Name": "Dionysus", "Score": 420},
		{"Name": "Hades", "Score": 666}
	]`

	t.Run("league from a Reader", func(t *testing.T) {
		database, dropDatabase := createTempFile(t, jsonLeague)
		defer dropDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Dionysus", 420},
			{"Hades", 666},
		}

		assertLeague(t, got, want)

		// test it can read the same league again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, dropDatabase := createTempFile(t, jsonLeague)
		defer dropDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Dionysus")
		want := 420

		assertScoreEquals(t, got, want)
	})

	t.Run("record player win", func(t *testing.T) {
		database, dropDatabase := createTempFile(t, jsonLeague)
		defer dropDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Hades")

		got := store.GetPlayerScore("Hades")
		want := 667
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, dropDatabase := createTempFile(t, jsonLeague)
		defer dropDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("New Player")

		got := store.GetPlayerScore("New Player")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with empty file", func(t *testing.T) {
		database, dropDatabase := createTempFile(t, "")
		defer dropDatabase()
		_, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
