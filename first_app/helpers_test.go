package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		log.Fatalf("didn't expect error but got %v", err)
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

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()

	got, err := NewLeague(body)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into Player slice, %v", body, err)
	}

	return got
}

func assertLeague(t testing.TB, got, wantedLeague []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague) {
		t.Errorf("got %v want %v", got, wantedLeague)
	}
}

func assertResponseBody(t testing.TB, got, want, name string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q for %q", got, want, name)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}
