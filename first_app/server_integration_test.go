package poker_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/Brett-Tanner/go_with_tests/firstapp"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, dropDatabase := poker.CreateTempFile(t, "[]")
	defer dropDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)
	poker.AssertNoError(t, err)

	server := poker.NewPlayerServer(store)
	player := "Preposterousa"
	wins := 3

	for i := 0; i < wins; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(t, player))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetScoreRequest(t, player))

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), fmt.Sprint(wins), player)
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewLeagueRequest())

		poker.AssertStatus(t, response.Code, http.StatusOK)

		got := poker.GetLeagueFromResponse(t, response.Body)
		want := []poker.Player{{player, wins}}
		poker.AssertLeague(t, got, want)
	})
}

func newPostWinRequest(t *testing.T, name string) *http.Request {
	t.Helper()

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}
