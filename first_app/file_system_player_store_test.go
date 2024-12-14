package main

import (
	"os"
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
		store := NewFileSystemPlayerStore(database)

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
		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Dionysus")
		want := 420

		assertScoreEquals(t, got, want)
	})

	t.Run("record player win", func(t *testing.T) {
		database, dropDatabase := createTempFile(t, jsonLeague)
		defer dropDatabase()
		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Hades")

		got := store.GetPlayerScore("Hades")
		want := 667
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, dropDatabase := createTempFile(t, jsonLeague)
		defer dropDatabase()
		store := NewFileSystemPlayerStore(database)

		store.RecordWin("New Player")

		got := store.GetPlayerScore("New Player")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create tempfile, %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
