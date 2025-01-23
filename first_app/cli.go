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
	PlayerPrompt   = "Please enter the number of players:\n"
	InputTypeError = "Please enter a number\n"
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

	userInput := c.readLine()
	c.game.Finish(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
