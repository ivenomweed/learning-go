package main

import "fmt"

// For loop

func main() {
	// Loop with single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// Alt Loop with single condition
	for i := range 3 {
		fmt.Println(i)
	}
	// Classic loop with initial/condition/after loop condition
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	// Infinite loop unless "break" out of it
	for {
		fmt.Println("loop")
		break
	}

	// use 'continue' to skip to next iteration
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
