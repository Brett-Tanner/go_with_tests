package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/Brett-Tanner/go_with_tests/firstapp"
)

var dummySpyAlerter = &SpyBlindAlerter{}

type SpyBlindAlerter struct {
	alerts []Alert
}

type Alert struct {
	scheduledAt time.Duration
	amount      int
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, Alert{duration, amount})
}

func TestCLI(t *testing.T) {
	t.Run("record Brett win from user input", func(t *testing.T) {
		want := "Brett"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", want))
		playerStore := &poker.StubPlayerStore{}

		cli := *poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, want)
	})

	t.Run("record Vika win from user input", func(t *testing.T) {
		want := "Vika"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", want))
		playerStore := &poker.StubPlayerStore{}

		cli := *poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, want)
	})

	t.Run(("scedules printing of blind values"), func(t *testing.T) {
		in := strings.NewReader("Vika wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})
}
