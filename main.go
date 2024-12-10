package main

import "fmt"

// Arrays

func main() {
	// Creates empty var with exactly n number of ints
	// by default ints are zero-valued
	var a [5]int
	fmt.Println("emp:", a)

	// set a value at an index
	a[4] = 100
	// get a value of an index
	fmt.Println("get:", a[4])
	// get length of an array
	fmt.Println("len:", len(a))

	// Initialize arryas with values
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// create a fixed length array automatically can also specify index for each value
	c := [...]int{100, 10: 400, 500}
	fmt.Println("idx:", c)

	// multi-dimentional arrays
	twoD := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println(twoD)
}
