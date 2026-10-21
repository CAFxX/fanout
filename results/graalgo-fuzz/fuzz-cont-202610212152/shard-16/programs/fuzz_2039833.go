package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]int)
    var v3 int16 = -32768
    v4 := make(map[uint32]uint32)
    v5 := GenS1{F0: 1}
    v3 += v3
    switch "a" {
    case "a":
        v1 = (v1 + 1)
    case "\n\t":
        v3 = -879
    case "hello":
        v6 := GenS1{F0: -268}
        _ = v6
    default:
        fmt.Println(v1, v3)
    }
    fmt.Println(v3)
    var v7 bool = (!(!false))
    v8 := GenS1{F0: 101}
    v9 := []int{83, -9223372036854775808, 127}
    var v10 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
