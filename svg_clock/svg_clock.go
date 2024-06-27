package svgclock

import (
	"fmt"
	"io"
	"math"
	"time"
)

var (
	clockCenter        float64 = 150
	hourHandLength     float64 = 50
	hoursInHalfClock   float64 = 6
	hoursInFullClock   float64 = 12
	minuteHandLength   float64 = 80
	minutesInHalfClock float64 = 30
	minutesInFullClock float64 = 60
	secondHandLength   float64 = 90
	secondsInHalfClock float64 = 30
	secondsInFullClock float64 = 60
)

type Point struct {
	X float64
	Y float64
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(secondHandPoint(t), secondHandLength)
	drawLine(w, p)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	drawLine(w, p)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(hourHandPoint(t), hourHandLength)
	drawLine(w, p)
}

func makeHand(p Point, l float64) Point {
	p = Point{p.X * l, p.Y * l}
	p = Point{p.X, -p.Y}
	return Point{p.X + clockCenter, p.Y + clockCenter}
}

func drawLine(w io.Writer, p Point) {
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(t.Second())))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / secondsInFullClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (math.Pi / (hoursInHalfClock / float64(t.Hour()%int(hoursInFullClock)))) + (minutesInRadians(t) / hoursInFullClock)
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
