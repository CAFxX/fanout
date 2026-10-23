package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 100
    var v2 bool = true
    var v3 int8 = (v1 | v1)
    v4 := make(map[string]string)
    var v5 uint = 42
    v2 = true
    fmt.Println(v3, v5)
    var v6 uint8 = 127
    switch "foo bar" {
    case "a":
        v2 = v2
    case "hello":
        var v7 uint8 = (v6 + (v6 / ((v6) | 1)))
        _ = v7
    default:
        var v8 uint32 = 245
        _ = v8
    }
    v5 += (v5 << ((v5) & 63))
    var v9 float32 = 1e10
    var v10 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
