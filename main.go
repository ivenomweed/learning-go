package main

import (
	"fmt"
	"maps"
)

func main() {
	// create an empty map
	m := make(map[string]int)

	// set key/value
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// get value based on key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// length of map
	fmt.Println("len:", len(m))

	// delete k/v
	delete(m, "k2")
	fmt.Println("map:", m)

	// clear all k/v
	clear(m)
	fmt.Println("map:", m)

	// missing keys and keys with zero values like 0 or "" check with optional second retur
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// same line map declaration
	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("map:", n)

	// maps std lib contains useful utilities
	n2 := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	if maps.Equal(n, n2) {
		fmt.Println(maps.Equal(n, n2))
	}

}
