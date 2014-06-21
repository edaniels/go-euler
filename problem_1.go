package main

import (
	"flag"
	"fmt"
)

func main() {

	var lim int
	flag.IntVar(&lim, "limit", 1000, "Sum of multiples of 3 or 5 from [1,limit)")
	flag.Parse()

	if lim < 0 {
		fmt.Printf("Limit (%d) cannot be negative. Exiting.\n", lim)
	}

	sum := 0

	for i := 1; i < lim; i++ {

		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Printf("Sum from 1 to %d: %d\n", lim, sum)
}
