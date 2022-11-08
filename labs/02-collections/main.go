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

	// map
	var m = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)

	// Struct
	// Defining Struct Type
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	// Declaring variable of struct type
	p := Person{"James", "Bond", 30}
	fmt.Println(p.FirstName)
	fmt.Println(p)
}
