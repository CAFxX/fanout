package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []uint32{}
    v2 := []int{0, 42}
    var v3 float64 = 3.14159
    v4 := GenS1{F0: 1, F1: 45}
    fmt.Println(v3)
    fmt.Println(v3)
    switch "hello" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
    default:
        _ = len(v2)
    }
    v3 = v3
    var v5 uint64 = 1
    if true {
        var v6 string = "a"
        v7 := GenS1{F0: (v5 - (v5 & (v5 % ((v5) | 1)))), F1: -1}
        switch v6 {
        case "\n\t":
            v5 += v5
        default:
            v2 = append(v2, 662)
        }
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: -549}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
