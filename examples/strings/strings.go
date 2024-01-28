package main

import (
	"fmt"
	"github.com/jon-ski/dry"
)

func main() {
	s := []string{"Hello", "Empty", "World"}
	fmt.Println(dry.StrJoinN(":", s...))
	// "Hello:Empty:World"

	// Change the "Empty" index to an empty string
	s = []string{"Hello", "", "World"}
	fmt.Println(dry.StrJoinN(":", s...))
	// "Hello:World"
}
