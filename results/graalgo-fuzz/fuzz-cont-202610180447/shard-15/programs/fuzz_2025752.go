package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int8 = -92
    var v2 uint32 = 4294967295
    v3 := []uint32{4294967295, (^(v2))}
    v4 := make(map[uint32]string)
    fmt.Println(v1)
    fmt.Println(v2)
    var v5 uint32 = 1
    v6 := GenS1{F0: -750}
    var v7 int = 9223372036854775807
    switch "foo bar" {
    case "\n\t":
        var v8 float64 = 3.14159
        _ = v8
    default:
        v9 := GenS1{F0: v7}
        _ = v9
    }
    v2 -= 42
    _ = len(v4)
    v2 = v2
    if true {
        v10 := GenS2{F0: (v1 >> ((255) & 63))}
        var v11 int = v7
        _ = v10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
