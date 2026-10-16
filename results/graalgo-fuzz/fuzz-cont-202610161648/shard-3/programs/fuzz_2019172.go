package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    v1 := []uint32{1, 4294967295, 121, 2}
    var v2 uint16 = 255
    var v3 int64 = 821
    var v4 int8 = -98
    fmt.Println(v3, v4, v3)
    v4 = (v4 | v4)
    var v5 int8 = v4
    fmt.Println(v5, v3)
    var v6 uint = 331
    switch "hello" {
    case "foo bar":
        v6 -= 42
    case "hello":
        fmt.Println(v6)
    default:
        v7 := GenS1{F0: 42, F1: v3}
        _ = v7
    }
    var v8 float64 = 1e10
    for fi9 := 0; fi9 < 8; fi9++ {
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        _ = fi9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
