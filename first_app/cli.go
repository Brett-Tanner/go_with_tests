package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	PlayerPrompt     = "Please enter the number of players:\n"
	InputTypeError   = "Please enter a number\n"
	WinnerInputError = "Invalid format for winner entry. Please try again\n"
)

type Game interface {
	Start(numPlayers int)
	Finish(winner string)
}

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

type ScheduledAlert struct {
	ScheduledAt time.Duration
	Amount      int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.ScheduledAt)
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numPlayers, err := strconv.Atoi(c.readLine())
	if err != nil {
		fmt.Fprint(c.out, InputTypeError)
		return
	}

	c.game.Start(numPlayers)

	winner, err := extractWinner(c.readLine())
	if err != nil {
		fmt.Fprint(c.out, WinnerInputError)
	}

	c.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	winnerSuffix := " wins"
	if !strings.Contains(userInput, winnerSuffix) {
		return "", fmt.Errorf("Winner entered incorrectly")
	}

	winner := strings.Replace(userInput, winnerSuffix, "", 1)

	return winner, nil
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
