package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 bool }
func genFn0(p0 int64, p1 int32, p2 int) uint8 {
    var v6 uint64 = 2
    p0 = -9223372036854775808
    v7 := []int{p2, (p2 / ((0) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 168
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]bool)
    var v3 int = 127
    v4 := []int{(v3 ^ v3)}
    var v5 float64 = 0.0
    var v8 int8 = -95
    var v9 uint32 = 265
    for fi10 := 0; fi10 < 6; fi10++ {
        if (0 != v8) {
            var v11 uint = 18446744073709551615
            _ = v11
        }
        v8 -= (v8 << ((290) & 63))
        var v12 int64 = 0
        _ = fi10
        _ = v12
    }
    if true {
        v4 = append(v4, (v3 << ((454) & 63)))
        v5 = 1.5
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                v5 -= v5
                _ = fi14
            }
            _ = fi13
        }
    }
    v5 += (v5 * v5)
    v15 := GenS1{F0: 0, F1: 42, F2: false}
    switch ("" + "hello") {
    case "\n\t":
        fmt.Println(v5)
    }
    v5 += 1e10
    v8 -= (v8 >> ((42) & 63))
    v9 += 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
