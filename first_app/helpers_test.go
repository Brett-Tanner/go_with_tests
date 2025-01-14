package poker

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		log.Fatalf("didn't expect error but got %v", err)
	}
}

func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
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

func NewLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func GetLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()

	got, err := NewLeague(body)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into Player slice, %v", body, err)
	}

	return got
}

func AssertLeague(t testing.TB, got, wantedLeague []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague) {
		t.Errorf("got %v want %v", got, wantedLeague)
	}
}

func AssertResponseBody(t testing.TB, got, want, name string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q for %q", got, want, name)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Errorf("got %d calls to RecordWin, wanted %d", len(store.WinCalls), 1)
	}
	if store.WinCalls[0] != winner {
		t.Errorf("expected %s to win but %s won", winner, store.WinCalls[0])
	}
}

func NewGetScoreRequest(t *testing.T, name string) *http.Request {
	t.Helper()

	path := fmt.Sprintf("/players/%v", name)
	response, _ := http.NewRequest(http.MethodGet, path, nil)
	return response
}
