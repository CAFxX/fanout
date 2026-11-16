package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint32) string {
    fmt.Println(p0)
    v5 := []int{-1}
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 = p0
        _ = fi6
    }
    _ = p0
    _ = v5
    return ("a" + "a")
}

func main() {
    var v1 float64 = float64((0.1 / (0.1 + 1e10)))
    var v2 int16 = 1000
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 18446744073709551615
    switch ("" + "hello") {
    case "\n\t":
        fmt.Println(v4)
    case "a":
        v4 -= v4
    case "hello":
        switch "hello" {
        case "hello":
            for fi7 := 0; fi7 < 13; fi7++ {
                v2 = -240
                _ = fi7
            }
        }
    default:
        fmt.Println(v4, v2, v4)
    }
    v4 = v4
    v8 := []int{}
    v9 := GenS1{F0: 18446744073709551615, F1: 233}
    var v10 int8 = 52
    v10 = v10
    switch "\n\t" {
    case "foo bar":
        var v11 int64 = -9223372036854775808
        _ = v11
    case "a":
        var v12 uint32 = 2
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint64 = 42
        var v15 uint8 = 128
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint64 = 4294967295
    fmt.Println(v1, v4, v4)
    v17 := GenS1{F0: (v4 ^ v16), F1: 42}
    var v18 uint8 = 128
    var v19 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
