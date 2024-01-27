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

