# dry - DRY Go Library for Common Functions

## Overview

`dry` is a Go library providing a collection of common generic functions for
slices and maps, designed to promote the DRY (Don't Repeat Yourself) principle
in Go programming. This library aims to reduce repetitive code patterns by
offering reusable, generic implementations for frequently needed operations.

## Installation

To use `dry` in your Go project, install it using `go get`:

```bash
go get github.com/jon-ski/dry
```

## Usage

`dry` provides generic functions for operations on slices and maps. Here's how
you can use these functions in your Go code:

### Examples

```go
import "github.com/jon-ski/dry"
```

#### Slices

```go
package main

import (
    "fmt"
    "github.com/jon-ski/dry"
)

func main() {
    intSlice := []int{1, 2, 3, 4, 5}
    fmt.Printf("Initial: %+v\n", intSlice)
    // Filter out even numbers
    intSlice = dry.Filter(intSlice, func(i int) bool {
        return i%2 != 0
    })
    fmt.Printf("Filtered: %+v\n", intSlice)
}
```

#### Maps

```go
package main

import (
    "fmt"
    "github.com/jon-ski/dry"
)

func main() {
    // Create a map[string]interface{} with some data
    m := map[string]interface{}{
        "name": "John Doe",
        "age":  42,
        "interests": "golang",
    }
    fmt.Printf("Initial: %+v\n", m)

    // Interests is a string, but we want a slice of strings
    // We can use the Map function to convert it
    m = dry.MapKeyFunc(m, "interests", func(i interface{}) interface{} {
        switch v := i.(type) {
        case string:
            return []string{v}
        case []string:
            return v
        default:
            return []string{}
        }
    })
    m["interests"] = append(m["interests"].([]string), "cycling")
    fmt.Printf("Updated: %+v\n", m)
}

```

## Function Documentation

Each function in `dry` is documented with comments explaining its purpose, parameters, and return values. For detailed documentation, see the source code or use `godoc`.

## License

`dry` is licensed under the [MIT License](LICENSE).

