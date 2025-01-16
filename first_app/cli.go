package poker

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type CLI struct {
	database PlayerStore
	in       *bufio.Scanner
	alerter  BlindAlerter
}

type ScheduledAlert struct {
	ScheduledAt time.Duration
	Amount      int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.ScheduledAt)
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		database: store,
		in:       bufio.NewScanner(in),
		alerter:  alerter,
	}
}

func (c *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func (c *CLI) PlayPoker() {
	c.scheduleBlindAlerts()
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
