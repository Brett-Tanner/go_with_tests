package main

import (
	"os"
	"time"

	svgclock "svg_clock"
)

func main() {
	t := time.Now()
	svgclock.SVGWriter(os.Stdout, t)
}
