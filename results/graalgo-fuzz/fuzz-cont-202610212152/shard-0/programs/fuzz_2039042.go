package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int16 }
func genFn0() uint {
    v6 := GenS1{F0: -1, F1: -1, F2: -1}
    var v7 float64 = 3.14159
    fmt.Println(v7)
    v8 := []int{1}
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []int{32767, 42}
    var v2 int = 42
    var v3 uint8 = 34
    v4 := make(map[string]uint32)
    var v5 float32 = (0.0 + 1.5)
    switch "a" {
    case "hello":
        switch "\n\t" {
        case "\n\t":
            v3 = (v3 / ((v3) | 1))
        }
    case "a":
        v9 := GenS1{F0: 0, F1: 188, F2: 32767}
        _ = v9
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v10 int16 = -909
    if true {
        v11 := GenS1{F0: -9223372036854775808, F1: 325, F2: v10}
        _ = v11
    } else {
        v5 -= v5
    }
    _ = len(v4)
    v2 += -490
    v3 = v3
    v2 = (v2 | (v2 + v2))
    var v12 int16 = (v10 ^ v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
