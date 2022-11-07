package main

import "fmt"

func main() {
	// array
	var x [5]int
	fmt.Println(x)

	// initialization
	x[0] = 100
	x[1] = 101
	x[3] = 103
	x[4] = 104
	fmt.Println(x)

	// Slices
	var s []int = x[1:4]
	fmt.Println(s)
}
