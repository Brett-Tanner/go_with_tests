package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	players := map[string]int{
		"Vika":  20,
		"Brett": 50,
	}
	store := StubPlayerStore{players, []string{}}
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
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func assertResponseBody(t testing.TB, got, want, name string) {
	if got != want {
		t.Errorf("got %q want %q for %q", got, want, name)
	}
}

func assertStatus(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}
