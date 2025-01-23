package poker

import "time"

type TexasHoldem struct {
	Database PlayerStore
	Alerter  BlindAlerter
}

func NewGame(store PlayerStore, alerter BlindAlerter) *TexasHoldem {
	return &TexasHoldem{store, alerter}
}

func (g *TexasHoldem) Start(numPlayers int) {
	blindIncrement := time.Duration(5+numPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}

	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.Alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (g *TexasHoldem) Finish(winner string) {
	g.Database.RecordWin(winner)
}
