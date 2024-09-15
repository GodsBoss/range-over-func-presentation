package main

import "fmt"

func main() {
	for n := range naturalNumbers {
		if n%2 == 0 {
			continue
		}
		if n > 20 {
			break
		}
		fmt.Println(n)
	}
}

func naturalNumbers(yield func(n int64) bool) {
	current := int64(0)
	for yield(current) {
		current++
	}
}
