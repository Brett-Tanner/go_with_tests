package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players:"

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game *Game
}

type ScheduledAlert struct {
	ScheduledAt time.Duration
	Amount      int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.ScheduledAt)
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: &Game{store, alerter},
	}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numPlayers, _ := strconv.Atoi(c.readLine())

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
