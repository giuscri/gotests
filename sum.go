package main

func Sum(xs []int) int {
	r := 0
	for _, x := range xs {
		r += x
	}
	return r
}

func SumAll(xxs ...[]int) []int {
	var r []int
	for _, xs := range xxs {
		r = append(r, Sum(xs))
	}
	return r
}
