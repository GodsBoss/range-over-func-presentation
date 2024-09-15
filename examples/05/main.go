package main

import "fmt"

func main() {
	yield := func(n int64) bool {
		if n%2 == 0 {
			return true
		}
		if n > 20 {
			return false
		}
		fmt.Println(n)
		return true
	}
	naturalNumbers(yield)
}

func naturalNumbers(yield func(n int64) bool) {
	current := int64(0)
	for yield(current) {
		current++
	}
}
