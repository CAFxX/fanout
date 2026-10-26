package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    var v1 string = "\n\t"
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    v4 := []uint32{}
    var v5 uint32 = 2
    var v6 float32 = 1.5
    if false {
        switch "" {
        case "hello":
            var v7 int64 = 9223372036854775807
            _ = v7
        case "a":
            v1 = (v1 + (v1 + v1))
        }
        fmt.Println(v5)
        v8 := v3[v5]
        _ = v8
    } else {
        v6 = float32(0.0)
        v1 = v1
        var v9 uint8 = 2
        _ = v9
    }
    v10 := GenS1{F0: -128, F1: 4294967295, F2: (128 == 9223372036854775807)}
    v11 := GenS1{F0: 0, F1: 18446744073709551615, F2: false}
    v5 = 2147483647
    v6 -= 0.1
    v6 = v6
    var v12 uint8 = 0
    var v13 uint = 215
    var v14 bool = false
    if (v1 < v1) {
        v15 := GenS1{F0: -1, F1: 11, F2: (-1 >= -538)}
        v12 = (v12 << ((65535) & 63))
        _ = v15
    }
    fmt.Println(v6)
    switch v1 {
    case "a":
        v12 = (v12 ^ 112)
    case "hello":
        var v16 float32 = (3.14159 * (-2.25 + (v6 + 3.14159)))
        _ = v16
    default:
        v17 := GenS1{F0: 42, F1: 4294967295, F2: (1.5 <= float64((v6 * (1e10 / float32((1.5 + 0.1))))))}
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
