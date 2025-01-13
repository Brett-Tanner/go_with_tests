package poker_test

import (
	"fmt"
	"strings"
	"testing"

	poker "github.com/Brett-Tanner/go_with_tests/firstapp"
)

func TestCLI(t *testing.T) {
	t.Run("record Brett win from user input", func(t *testing.T) {
		want := "Brett"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", want))
		playerStore := &poker.StubPlayerStore{}

		cli := *poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, want)
	})

	t.Run("record Vika win from user input", func(t *testing.T) {
		want := "Vika"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", want))
		playerStore := &poker.StubPlayerStore{}

		cli := *poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, want)
	})
}
