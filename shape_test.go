package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := r.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100},
		{name: "Circle", shape: Circle{Radius: 10.0}, want: 31.41592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want
			if got != want {
				t.Errorf("%#v: got %.2f want %.2f", tt.shape, got, want)
			}
		})
	}
}
