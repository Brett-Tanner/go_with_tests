package main

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name:    "Rectangle",
			shape:   Rectangle{Width: 10.0, Height: 10.0},
			hasArea: 100.0,
		},
		{
			name:    "Circle",
			shape:   Circle{Radius: 10.0},
			hasArea: 314.1592653589793,
		},
		{
			name:    "Triangle",
			shape:   Triangle{Base: 12.0, Height: 6.0},
			hasArea: 36.0,
		},
	}

	for _, args := range areaTests {
		t.Run(args.name, func(t *testing.T) {
			got := args.shape.Area()
			if got != args.hasArea {
				t.Errorf("%#v got %g want %g", args.shape, got, args.hasArea)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	got := Rectangle{10.0, 10.0}.Perimeter()
	want := 40.0
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
