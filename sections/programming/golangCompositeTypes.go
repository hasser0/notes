package main

import "fmt"

func main() {
	// Arrays
	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := [...]string{"Hello", "Entire", "World"}
	a3 := [5][2]float{}

	// Slice
	s1 := []int{1, 2, 3, 4, 5}
	s2Length := 10
	s2Capacity := 20
	s2 := make([]string, s2Length, s2Capacity)
	s3 := [][]float{[]float{3.1416, 2.7172}, []float{9.81}}

	xs = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Slicing a slice
	xs[1:4]
	xs[5:]
	xs[:9]
	// Append to slice
	xs = append(xs, 10, 11, 12)
	// Delete from slice
	xs = append(xs[0:4], xs[5:]...)
	// Copy slice
	xsNew := make([]int, len(xs), len(x2) * 2)
	copy(xsNew, xs)

}