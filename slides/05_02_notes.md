[Previous: How do they work?](./05_01_translation.md)

# Notes on "How do they work?"

## Calling `yield` after it returned `false`

Our custom `yield` function happily works when called again after it returned `false`. When doing this with proper iterators, a `panic: runtime error: range function continued iteration after function for loop body returned false` is thrown.

## `return'

A `for` loop may be ended by `return`ing values, exiting the function that contains the loop. Our example did not cover that case. We could do this, but it would be unnecessarely complex. Suffice to say that `return` works as expected from a `for` loop.

## `panic`

Throwing a `panic` works as expected.

[Next: How do we implement them?](./06.md)
