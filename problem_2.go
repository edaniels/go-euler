package main

import (
	"flag"
	"fmt"
)

func fibonacci() func() int {

	a, b := 0, 1

	return func() int {

		val := a
		a, b = b, a+b
		return val
	}
}

func main() {

	var lim int
	flag.IntVar(&lim, "limit", 4000000, "Limit of Fibonacci values")
	flag.Parse()

	if lim < 0 {
		fmt.Printf("Limit (%d) cannot be negative. Exiting.\n", lim)
	}

	fibGen := fibonacci()
	sum := 0

	for currVal := fibGen(); currVal <= lim; currVal = fibGen() {

		if currVal%2 == 0 {
			sum += currVal
		}
	}

	fmt.Printf("Sum of even Fibonacci numbers [0, %d]: %d\n", lim, sum)
}
