package svgclock

import (
	"math"
	"time"
)

var (
	CENTER             float64 = 150
	SECOND_HAND_LENGTH float64 = 90
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * SECOND_HAND_LENGTH, p.Y * SECOND_HAND_LENGTH}
	p = Point{p.X, -p.Y}
	p = Point{p.X + CENTER, p.Y + CENTER}
	return p
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}
