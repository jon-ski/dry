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
