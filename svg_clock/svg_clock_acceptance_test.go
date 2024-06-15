package svgclock

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Lines   []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriter(t *testing.T) {
	cases := []struct {
		Time time.Time
		Line Line
	}{
		{simpleTime(1, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}

	for _, c := range cases {
		b := bytes.Buffer{}
		SVGWriter(&b, c.Time)

		svg := SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		if !containsLine(c.Line, svg.Lines) {
			t.Errorf("Expected secondhand line %+v in SVG lines %+v", c.Line, svg.Lines)
		}
	}
}

func containsLine(l Line, lines []Line) bool {
	for _, line := range lines {
		if line == l {
			return true
		}
	}

	return false
}
