package main

import (
	"strings"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("league from a Reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Dionysus", "Score": 420},
			{"Name": "Hades", "Score": 666}
		]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Dionysus", 420},
			{"Hades", 666},
		}

		assertLeague(t, got, want)
	})
}
