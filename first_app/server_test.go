package poker_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"

	poker "github.com/Brett-Tanner/go_with_tests/firstapp"
)

func TestGETPlayers(t *testing.T) {
	players := map[string]int{
		"Vika":  20,
		"Brett": 50,
	}
	store := poker.StubPlayerStore{players, []string{}, nil}
	server := poker.EnsurePlayerServerCreated(t, &store)

	t.Run("returns score for requested player", func(t *testing.T) {
		for name, score := range players {
			response := httptest.NewRecorder()
			request := poker.NewGetScoreRequest(t, name)

			server.ServeHTTP(response, request)

			poker.AssertStatus(t, response.Code, http.StatusOK)
			poker.AssertResponseBody(t, response.Body.String(), fmt.Sprint(score), name)
		}
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := poker.NewGetScoreRequest(t, "Artemis")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{},
		[]string{},
		nil,
	}
	server := poker.EnsurePlayerServerCreated(t, &store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusAccepted)
		poker.AssertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []poker.Player{
		{"Vika", 20},
		{"Brett", 21},
		{"Zagreus", 50},
	}

	store := poker.StubPlayerStore{nil, nil, wantedLeague}
	server := poker.EnsurePlayerServerCreated(t, &store)

	t.Run("returns the league table as JSON", func(t *testing.T) {
		request := poker.NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := poker.GetLeagueFromResponse(t, response.Body)

		poker.AssertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, poker.JsonContentType)
		poker.AssertLeague(t, got, wantedLeague)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := poker.EnsurePlayerServerCreated(t, &poker.StubPlayerStore{})

		request := poker.NewGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("when we get a message over a websocket it is the winner of a game", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		winner := "Vika"
		playerServer := poker.EnsurePlayerServerCreated(t, store)

		server := httptest.NewServer(playerServer)
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws := ensureWebSocketDialed(t, wsURL)
		defer ws.Close()
		writeWebSocketMessage(t, ws, winner)

		time.Sleep(100 * time.Millisecond)
		poker.AssertPlayerWin(t, store, winner)
	})
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type %s, got %v", want, response.Result().Header)
	}
}

func ensureWebSocketDialed(t *testing.T, url string) *websocket.Conn {
	t.Helper()

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a WebSocket connection on %s %v", url, err)
	}

	return ws
}

func writeWebSocketMessage(t *testing.T, conn *websocket.Conn, message string) {
	t.Helper()

	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over WebSocket connection %v", err)
	}
}
