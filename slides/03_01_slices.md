[Previous: Can't we use existing mechanisms?](./03.md)

# Can't we use existing mechanisms?

## Slices

Create a slice and iterate over that. What are the drawbacks?

* Parts of a slice may have been created for nothing if only a part of that slice is used in the end.
* Can't start "doing work" before the slice is constructed.
* You cannot model a potentially infinite amount of values.
* Not really suited when processing elements would add elements to process.
* Functions for filtering or mapping values create intermediate slices.
* Iteration is limited to `int, [V any]` pairs.

Arrays share these traits.

[Next: Can't we use maps?](./03_02_maps.md)
