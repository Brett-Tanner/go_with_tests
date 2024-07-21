package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func TestGETPlayers(t *testing.T) {
	players := map[string]int{
		"Vika":  20,
		"Brett": 50,
	}
	store := StubPlayerStore{players}
	server := &PlayerServer{store: &store}

	t.Run("returns score for requested player", func(t *testing.T) {
		for name, score := range players {
			response := httptest.NewRecorder()
			request, err := newGetScoreRequest(name)
			if err != nil {
				t.Fatal("Request failed")
			}

			server.ServeHTTP(response, request)

			assertResponseBody(t, response.Body.String(), fmt.Sprint(score), name)
		}
	})
}

func newGetScoreRequest(name string) (*http.Request, error) {
	path := fmt.Sprintf("/players/%v", name)
	return http.NewRequest(http.MethodGet, path, nil)
}

func assertResponseBody(t testing.TB, got, want, name string) {
	if got != want {
		t.Errorf("got %q want %q for %q", got, want, name)
	}
}
