package main

import "testing"

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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		checkArea(t, r, 100.0)
	})

	t.Run("circle", func(t *testing.T) {
		c := Circle{10.0}
		checkArea(t, c, 31.41592653589793)
	})
}
