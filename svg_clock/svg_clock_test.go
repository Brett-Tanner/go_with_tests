package svgclock

import (
	"fmt"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	cases := []struct {
		Time  time.Time
		Point Point
	}{
		{
			time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC),
			Point{X: CENTER, Y: CENTER - 90},
		},
		{
			time.Date(1337, time.January, 1, 0, 30, 0, 0, time.UTC),
			Point{X: CENTER, Y: CENTER + 90},
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("second hand points at %v at %q", test.Point, test.Time.Format("3:04PM")), func(t *testing.T) {
			got := SecondHand(test.Time)

			if got != test.Point {
				t.Errorf("Got %v want %v", got, test.Point)
			}
		})
	}
}
