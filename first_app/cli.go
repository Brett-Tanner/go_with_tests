package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	database PlayerStore
	in       io.Reader
}

func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.in)
	reader.Scan()
	c.database.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
