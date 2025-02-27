[Previous: How do we use them?](./04.md)

# How do they work?

Now that we have seen how to use an iterator, let's examine how it works. Basically, the `for` loop is converted into a `yield` function and passed to the iterator. Let's do that ourselves. Here's the relevant part of the previous example:

```go
	for n := range naturalNumbers {
		if n%2 == 0 {
			continue
		}
		if n > 20 {
			break
		}
		fmt.Println(n)
	}
```

Converting the loop results in this code, which is invalid (`continue` and `break` without a loop, no return values):

```go
	yield := func(n int64) bool {
		if n%2 == 0 {
			continue
		}
		if n > 20 {
			break
		}
		fmt.Println(n)
	}
```

Note that the function signature `func(int64) bool` is the same as the type for `naturalNumbers`' `yield` argument.

Now let's make this valid Go code. Replace `continue` with `return true`, `break` with `return false` and default to `return true` at the end:

```go
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
```

Looks like a valid function. Let's pass it to `naturalNumbers` and look what happens. Here is the full example:

```go
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
```

Output ([example](../examples/05/main.go)):
```
1
3
5
7
9
11
13
15
17
19
```

It's the same output. This is roughly what happens under the hood when ranging over a function.

[Next: Notes](./05_02_notes.md)
