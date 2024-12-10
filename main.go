package main

import "fmt"

// Variables

func main() {
	// Declare one or more variable with var
	var a = "first"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// Variable type is infered if variable is initialized with value
	// Variable initialized without value needs to have a type declaration and will have the types coresponding zero-value
	var d = true
	fmt.Println(d)

	var e bool
	fmt.Println(e)

	// Short hand for declaring and initializing a variable
	f := "appple"
	fmt.Println(f)

	g, h := 69, false
	fmt.Println(g, h)
}
