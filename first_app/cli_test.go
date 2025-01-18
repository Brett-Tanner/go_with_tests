package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/Brett-Tanner/go_with_tests/firstapp"
)

var dummyPlayerStore = &poker.StubPlayerStore{}

type SpyBlindAlerter struct {
	alerts []poker.ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, poker.ScheduledAlert{duration, amount})
}

func TestCLI(t *testing.T) {
	t.Run("record Brett win from user input", func(t *testing.T) {
		dummySpyAlerter := &SpyBlindAlerter{}
		stdout := &bytes.Buffer{}
		want := "Brett"
		in := strings.NewReader(fmt.Sprintf("5\n%s wins\n", want))

		cli := *poker.NewCLI(dummyPlayerStore, in, stdout, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, dummyPlayerStore, want)
	})

	t.Run("prompts the user to enter the number of players", func(t *testing.T) {
		dummySpyAlerter := &SpyBlindAlerter{}
		in := strings.NewReader("7\n")
		stdout := &bytes.Buffer{}
		cli := poker.NewCLI(dummyPlayerStore, in, stdout, dummySpyAlerter)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(dummySpyAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, dummySpyAlerter.alerts)
				}

				got := dummySpyAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("record Vika win from user input", func(t *testing.T) {
		dummySpyAlerter := &SpyBlindAlerter{}
		stdout := &bytes.Buffer{}
		want := "Vika"
		in := strings.NewReader(fmt.Sprintf("5\n%s wins\n", want))
		playerStore := &poker.StubPlayerStore{}

		cli := *poker.NewCLI(playerStore, in, stdout, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, want)
	})

	t.Run(("schedules printing of blind values"), func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("5\nVika wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, stdout, blindAlerter)
		cli.PlayPoker()

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

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= 1 {
					t.Fatalf("alert %d was not scheduled, %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
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
