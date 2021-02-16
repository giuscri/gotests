package main

type Rectangle struct {
	w float64
	h float64
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.w + 2*r.h
}

func (r *Rectangle) Area() float64 {
	return r.w * r.h
}
