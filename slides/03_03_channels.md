[Previous: Can't we use maps?](./03_02_maps.md)

# Can't we use existing mechanisms?

## Channels

What are the drawbacks of using channels?

* Synchronization mechanisms lead to a bit of overhead.
* Functions for filtering or mapping channel values typically involve additional goroutines, leading to more overhead.
* When iteration ends prematurely (`break`), producers need to be closed to avoid resource leakage ([problem example](../examples/03_03/main.go)).
* Iteration is constrained to single `[V any]` values.

[Next: How do we use them?](./04.md)
