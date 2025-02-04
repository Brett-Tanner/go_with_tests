package poker_test

import (
	"fmt"
	"io"
	"testing"
	"time"

	poker "github.com/Brett-Tanner/go_with_tests/firstapp"
)

var dummyPlayerStore = &poker.StubPlayerStore{}

func TestGameStart(t *testing.T) {
	t.Run("prompts the user to enter the number of players", func(t *testing.T) {
		spyAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(dummyPlayerStore, spyAlerter)

		game.Start(7, io.Discard)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkSchedulingCases(t, cases, *spyAlerter)
	})

	t.Run(("schedules printing of blind values"), func(t *testing.T) {
		spyAlerter := &SpyBlindAlerter{}
		game := &poker.TexasHoldem{dummyPlayerStore, spyAlerter}

		game.Start(5, io.Discard)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(t, cases, *spyAlerter)
	})
}

func checkSchedulingCases(t *testing.T, cases []poker.ScheduledAlert, alerter SpyBlindAlerter) {
	t.Helper()

	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(alerter.alerts) <= 1 {
				t.Fatalf("alert %d was not scheduled, %v", i, alerter.alerts)
			}

			got := alerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}

func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	t.Helper()

	gotAmount := got.Amount
	if gotAmount != want.Amount {
		t.Errorf("got amount %d want amount %d", gotAmount, want.Amount)
	}

	gotScheduleTime := got.ScheduledAt
	if gotScheduleTime != want.ScheduledAt {
		t.Errorf("got scheduled time of %v, want %v", gotScheduleTime, want.ScheduledAt)
	}
}

func TestGameFinish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewGame(store, &SpyBlindAlerter{})
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}
