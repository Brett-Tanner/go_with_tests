package svgclock

import "time"

var CENTER float64 = 150

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
