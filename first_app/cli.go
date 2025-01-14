package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	database PlayerStore
	in       *bufio.Scanner
	alerter  BlindAlerter
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		database: store,
		in:       bufio.NewScanner(in),
		alerter:  alerter,
	}
}

func (c *CLI) PlayPoker() {
	c.alerter.ScheduleAlertAt(5*time.Second, 100)
	userInput := c.readLine()
	c.database.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
