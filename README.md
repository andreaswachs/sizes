# Sizes

`sizes` is a Golang package that attempts to get rid of anonymous size calculations in projects, where there is poor communication between developers in terms of what the unit size of the calculation attempts to be (bytes, kilobytes,...). 

This library also attempts to get rid of error prone calculations by using the builder pattern to facilitate simple size calculations. You "build" the results by finishing the function *chains* with either `Calculate()` or `CalculateAs(unit)`.

Shorthand functions have been created to reduce verbosity where you might not care to start each calculation by instantiation a `SizeBuilder`. You just start by choosing the unit that you want to start calculating on and specify a multiplier and then calculate. See the examples section for more.

## Caveats

Overflow risks: If you're computing the byte size of many exabytes, you might run into overflow issues, since the calculations return `uint64` and might run out of space for the real number.

Wrong use case: If you're doing many calculations where you compute some byte sizes, it might be faster for you to do the calculations by hand, as a builder pattern is used here and will create a struct each time you calculate a size.

## Examples

Calculating 100 megabytes in bytes:

```go
sizes.Megabytes().Multiply(100).Calculate()
```

Calculating 2 million bytes as kilobytes (please be aware of rounding down since dealing with integers):

```go
sizes.Bytes().Multiply(2_000_000).CalculateAs(sizes.Kilobytes)
```
