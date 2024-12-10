package main

import (
	"fmt"
	"slices"
)

func main() {
	// empty string slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// empty silce with non-zero value
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// set values by index
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	// Append to slice
	s = append(s, "d")
	fmt.Println(s)
	s = append(s, "e", "f")
	fmt.Println(s)

	// copy slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice operator
	l := s[2:5]
	fmt.Println(l)

	// slices upto (but excluding)
	l = s[:5]
	fmt.Println(l)

	// slices from (and including)
	l = s[2:]
	fmt.Println(l)

	// single line slice declaration
	t := []string{}
	fmt.Println(t)

	t = append(t, "a")
	fmt.Println(t)

	t2 := make([]string, len(t))
	copy(t2, t)
	fmt.Println(t, t2)

	// slices std lib containes useful utilities
	if slices.Equal(t, t2) {
		fmt.Println(slices.Equal(t, t2))
	}
}
