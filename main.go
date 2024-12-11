package main

import "fmt"

// Pointers
func zeroVal(ival int) {
	ival = 0
}
func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// value of i doesn't change
	zeroVal(i)
	fmt.Println("zeroval:", i)

	// value of i changes because the pointer to the memory address
	//  is being passed instead of a
	// reference so when zeroPtr changes the value it changes the value
	// in the memory address directly
	zeroPtr(&i)
	fmt.Println("zeroptr:", i)

	// Memory address
	fmt.Println("pointer", &i)
}
