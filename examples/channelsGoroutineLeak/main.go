package main

import (
	"fmt"
	"runtime"
)

func main() {
	printGoroutines() // Count is 1, as there's only main().

	for n := range numbers(4) {
		if n > 4 {
			break
		}
		fmt.Println(n)
	}

	printGoroutines() // Count is 1, as goroutine spawned by numbers(4) exited.

	for n := range numbers(8) {
		if n > 4 {
			break
		}
		fmt.Println(n)
	}

	printGoroutines() // Count is 2, as goroutine spawned by numbers(8) never exits. Resource leak!
}

func printGoroutines() {
	fmt.Printf("Goroutine count: %d\n", runtime.NumGoroutine())
}

func numbers(max int) <-chan int {
	ch := make(chan int)
	n := 0
	go func() {
		for n < max {
			ch <- n
			n++
		}
		close(ch)
	}()
	return ch
}
