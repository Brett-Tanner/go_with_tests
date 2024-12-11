package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	players := map[string]int{
		"Vika":  20,
		"Brett": 50,
	}
	store := StubPlayerStore{players, []string{}, nil}
	server := NewPlayerServer(&store)

	t.Run("returns score for requested player", func(t *testing.T) {
		for name, score := range players {
			response := httptest.NewRecorder()
			request := newGetScoreRequest(name)

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, http.StatusOK)
			assertResponseBody(t, response.Body.String(), fmt.Sprint(score), name)
		}
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Artemis")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func newGetScoreRequest(name string) *http.Request {
	path := fmt.Sprintf("/players/%v", name)
	response, _ := http.NewRequest(http.MethodGet, path, nil)
	return response
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, wanted %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("expected %s to win but %s won", player, store.winCalls[0])
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Vika", 20},
		{"Brett", 21},
		{"Zagreus", 50},
	}

	store := StubPlayerStore{nil, nil, wantedLeague}
	server := NewPlayerServer(&store)

	t.Run("returns the league table as JSON", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		assertLeague(t, got, wantedLeague)
	})
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

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type %s, got %v", want, response.Result().Header)
	}
}
