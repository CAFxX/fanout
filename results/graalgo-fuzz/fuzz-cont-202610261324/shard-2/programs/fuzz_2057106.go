package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 0
    v3 := []string{"\n\t", "\n\t", "\n\t"}
    v2 = (^(v2))
    switch 2147483647 {
    case -511:
        v2 = 1
    default:
        v1[""] = 42
    }
    v2 = v2
    var v4 uint = 0
    fmt.Println(v2, v2)
    v5 := GenS1{F0: 42}
    v6 := GenS1{F0: 42}
    v4 -= v4
    var v7 uint16 = 42
    var v8 int8 = 100
    var v9 float32 = -2.25
    fmt.Println(v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
