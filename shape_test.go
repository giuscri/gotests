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
	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := r.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
