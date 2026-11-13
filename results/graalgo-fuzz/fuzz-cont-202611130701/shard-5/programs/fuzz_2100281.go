package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 string = ("a" + "")
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    v2 = v2
    v2 = ("\n\t" + v2)
    var v5 uint64 = 995
    switch "a" {
    case "\n\t":
        v6 := GenS1{F0: (!false), F1: (^(v5))}
        _ = v6
    }
    var v7 int32 = 103
    v8 := GenS2{F0: -32768, F1: 1, F2: v7}
    switch 0 {
    case 887:
        var v9 float32 = (3.14159 + float32((-2.25 / 1e10)))
        _ = v9
    default:
        v2 = v2
    }
    v5 += v5
    var v10 uint32 = 0
    v11 := GenS2{F0: -863, F1: 42, F2: (v7 - 1)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
