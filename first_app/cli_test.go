package poker

import (
	"fmt"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Brett win from user input", func(t *testing.T) {
		want := "Brett"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", want))
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, want)
	})

	t.Run("record Vika win from user input", func(t *testing.T) {
		want := "Vika"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", want))
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, want)
	})
}
