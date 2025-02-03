package poker_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	poker "github.com/Brett-Tanner/go_with_tests/firstapp"
)

type SpyBlindAlerter struct {
	alerts []poker.ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, poker.ScheduledAlert{duration, amount})
}

type GameSpy struct {
	BlindAlert  []byte
	StartCalled bool
	StartedWith int

	FinishedCalled bool
	FinishedWith   string
}

func (g *GameSpy) Start(numPlayers int, to io.Writer) {
	g.StartCalled = true
	g.StartedWith = numPlayers
	to.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("record Brett win from user input", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		want := "Brett"
		in := strings.NewReader(fmt.Sprintf("5\n%s wins\n", want))

		cli := *poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertCLIWin(t, game.FinishedWith, want)
	})

	t.Run("record Vika win from user input", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		want := "Vika"
		in := strings.NewReader(fmt.Sprintf("5\n%s wins\n", want))

		cli := *poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertCLIWin(t, game.FinishedWith, want)
	})

	t.Run("when invalid winner string entered, prompts again with error", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		in := strings.NewReader("3\nBrett is da coolest")

		cli := *poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(
			t, stdout,
			poker.PlayerPrompt+poker.WinnerInputError,
		)
	})

	t.Run("when non-numeric value entered, prints error and doesn't start game", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		in := strings.NewReader("pizza")

		cli := *poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(
			t, stdout,
			poker.PlayerPrompt+poker.InputTypeError,
		)
	})
}

func assertCLIWin(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("expected %q to win but %q won", want, got)
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Error("game started but shouldn't have")
	}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
