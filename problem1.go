package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	ch := make(chan int)
	sum := 0

	for i := 1; i < 1000; i++ {
		go multThreeFive(i, ch)
	}

	for i := 1; i < 1000; i++ {
		sum += <-ch
	}

	fmt.Printf("The sum of all the multiples of 3 or 5 below 1000 is %d.\n", sum)

	end := time.Since(start).Milliseconds()
	fmt.Printf("(execution time: %dms)\n", end)
}

func multThreeFive(i int, ch chan<- int) {

	if i%3 == 0 || i%5 == 0 {
		ch <- i
	} else {
		ch <- 0
	}

}
