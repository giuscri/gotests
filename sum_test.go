package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// for i, v := range got {
	// 	if i >= len(want) {
	// 		t.Errorf("got longer slice than wanted: %q, %q", got, want)
	// 	}
	// 	if v != want[i] {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// }
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
