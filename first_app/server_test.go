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

var dummyGame = &GameSpy{}

func TestGETPlayers(t *testing.T) {
	players := map[string]int{
		"Vika":  20,
		"Brett": 50,
	}
	store := poker.StubPlayerStore{players, []string{}, nil}
	server := poker.EnsurePlayerServerCreated(t, &store, dummyGame)

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
	server := poker.EnsurePlayerServerCreated(t, &store, dummyGame)

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
	server := poker.EnsurePlayerServerCreated(t, &store, dummyGame)

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
		server := poker.EnsurePlayerServerCreated(t, &poker.StubPlayerStore{}, dummyGame)

		request := poker.NewGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("start a game with 3 players and declare Vika the winner", func(t *testing.T) {
		game := &GameSpy{}
		winner := "Vika"
		store := &poker.StubPlayerStore{}

		server := httptest.NewServer(poker.EnsurePlayerServerCreated(t, store, game))
		ws := ensureWebSocketDialed(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")
		defer ws.Close()
		defer server.Close()

		writeWebSocketMessage(t, ws, "3")
		writeWebSocketMessage(t, ws, winner)

		time.Sleep(100 * time.Millisecond)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
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

func assertGameStartedWith(t testing.TB, game *GameSpy, players int) {
	t.Helper()

	if game.StartedWith != players {
		t.Errorf("wanted Start called with %d but got %d", players, game.StartedWith)
	}
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()

	if game.FinishedWith != winner {
		t.Errorf("expected %s to win but got %s", winner, game.FinishedWith)
	}
}
