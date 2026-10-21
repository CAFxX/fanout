package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int64, p1 int8) uint {
    var v4 float64 = float64(1e10)
    p1 = (p1 | (-(p1)))
    v5 := GenS1{F0: 42, F1: 122}
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 42
    var v2 bool = (-1 >= 963)
    v3 := []string{"foo bar", "", ""}
    var v7 bool = true
    var v8 uint = genFn0(1, 43)
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v7 = v7
    _ = len(v3)
    switch "a" {
    case "foo bar":
        v2 = false
    default:
        fmt.Println(v7)
    }
    v1 = (v1 << (((^(v8))) & 63))
    var v9 int64 = -812
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(fi10, v9, v7)
        var v11 int8 = 42
        _ = fi10
        _ = v11
    }
    var v12 int8 = 1
    fmt.Println(v7, v7)
    v1 += (^(v1))
    v13 := GenS1{F0: (v1 & v1), F1: 42}
    v2 = false
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
