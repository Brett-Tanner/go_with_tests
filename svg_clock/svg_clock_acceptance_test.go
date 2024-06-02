package svgclock

import (
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	cases := []struct {
		Time  time.Time
		Point Point
	}{
		{
			simpleTime(0, 0, 0), Point{X: CENTER, Y: CENTER - 90},
		},
		{
			simpleTime(0, 0, 30), Point{X: CENTER, Y: CENTER + 90},
		},
	}

	for _, test := range cases {
		t.Run(testName(test.Time), func(t *testing.T) {
			got := SecondHand(test.Time)

			if got != test.Point {
				t.Errorf("Got %v want %v", got, test.Point)
			}
		})
	}
}
