package main

import "fmt"

// Functions

// function to add two integers
func plus(a int, b int) int {
	return a + b
}

// function to add three integers
func plusPlus(a, b, c int) int {
	return a + b + c
}

// function that returns multiple values
func vals() (int, int) {
	return 3, 7
}

// variadic function
func sum(nums ...int) {
	fmt.Print(nums, "")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// closures with annonymous fucntions
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursive functions
func fact(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println(n)
	return n * fact(n-1)
}

func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(plusPlus(1, 2, 3))
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	sum(1, 2, 3, 4, 5)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(fact(7))
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
