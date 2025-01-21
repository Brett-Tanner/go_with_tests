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
}
