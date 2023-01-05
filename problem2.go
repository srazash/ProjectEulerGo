package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	fmt.Printf("The sum of the even values from a Fibonacci sequence not exceeding 4,000,000 is %d.\n", sumEven(fibonacci(4000000)))

	end := time.Since(start).Milliseconds()
	fmt.Printf("(execution time: %dms)\n", end)
}

func fibonacci(t int) []int {

	fib := []int{}

	i := 1
	j := 2

	fib = append(fib, i)
	fib = append(fib, j)

	for j < t {

		x := i + j
		fib = append(fib, x)
		i = j
		j = x

	}

	return fib
}

func sumEven(fib []int) int {

	sum := 0

	for _, e := range fib {
		if e%2 == 0 {
			sum += e
		}
	}

	return sum

}
