[![Unit tests](https://github.com/andreaswachs/sizes/actions/workflows/unit-tests.yml/badge.svg)](https://github.com/andreaswachs/sizes/actions/workflows/unit-tests.yml)
[![pkg.go.dev](https://img.shields.io/badge/DOCS-pkg.go.dev-informational)](https://pkg.go.dev/github.com/andreaswachs/sizes)

# Sizes

`sizes` is a Golang package that attempts to produce byte sizes that are obvious for the reader, instead of having nondescript calculations and depending on developers to comment what they were calculating.


## Caveats

Overflow risks: If you're computing the byte size of many exabytes, you might run into overflow issues, since the calculations return `uint64` and might run out of space for the real number.

Wrong use case: If you're doing many calculations where you compute some byte sizes, it might be faster for you to do the calculations by hand, as a builder pattern is used here and will create a struct each time you calculate a size.

## Examples

Calculating 100 megabytes in bytes:

```go
sizes.Megabytes(100)
```

Bytes are the default unit that the sizes are returned in. If you wish to calculate 100 megabytes in kilobytes:

```go
sizes.MegabytesAs(100, sizes.Kilobyte)
```

You can also go the other way, calculating from lower units to larger.

Calculating 2 million bytes as kilobytes (please be aware of rounding down since dealing with integers):

```go
sizes.BytesAs(2_000_000, sizes.Megabyte)
```
