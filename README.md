# wildcard [![Test](https://github.com/vikpe/wildcard/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/vikpe/wildcard/actions/workflows/test.yml) [![codecov](https://codecov.io/gh/vikpe/wildcard/branch/main/graph/badge.svg)](https://codecov.io/gh/vikpe/wildcard)

> Wildcard pattern matching in Go (Golang)

## Install

```shell
go get github.com/vikpe/wildcard
```

## Usage

```go
import "github.com/vikpe/wildcard"

// strings
wildcard.Match("foo*", "foobar")   // true

wildcard.MatchCI("FOO*", "foobar") // true
wildcard.MatchCI("foo*", "FOOBAR") // true

// slice of strings
wildcard.MatchSlice("foo*", []string{"foobar", "barfoo"})   // true
wildcard.MatchSlice("*foo", []string{"foobar", "barfoo"})   // true

wildcard.MatchSliceCI("FOO*", []string{"foobar", "barfoo"})   // true
wildcard.MatchSliceCI("*foo", []string{"FOOBAR", "BARFOO"})   // true

```

## Functions

### Match

> String wildcard pattern match (case sensitive)

```go
wildcard.Match(pattern string, haystack string) bool
```

### MatchCI

> String wildcard pattern match (case insensitive)

```go
wildcard.MatchCI(pattern string, haystack string) bool
``` 

### MatchSlice

> Slice of strings wildcard pattern match (case sensitive)

```go
wildcard.MatchSlice(pattern string, haystack []string) bool
``` 

### MatchSliceCI

> Slice of strings wildcard pattern match (case insensitive)

```go
wildcard.MatchSliceCI(pattern string, haystack []string) bool
``` 

## Benchmarks

```shell
go test ./... -benchmem -bench=.

cpu: AMD Ryzen 5 5600X 6-Core Processor             
Match/no_wildcards-12          206281431       5.757 ns/op       0 B/op       0 allocs/op
Match/single_wildcard-12       22629099        50.38 ns/op       0 B/op       0 allocs/op
Match/multiple_wildcards-12    14568752        69.47 ns/op       0 B/op       0 allocs/op
MatchCI/no_wildcards-12        28999945        38.74 ns/op       5 B/op       1 allocs/op
MatchCI/single_wildcard-12     12787585        91.05 ns/op       8 B/op       1 allocs/op
MatchCI/multiple_wildcards-12  11541874        101.4 ns/op       5 B/op       1 allocs/op
```

## See also

* [**go-wildcard**](https://github.com/IGLOU-EU/go-wildcard) - Supports both `*` and `?` wildcard patterns.
