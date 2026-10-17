package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int; F1 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 42
    var v3 int = -9223372036854775808
    if true {
        v4 := []int{}
        var v5 float64 = 0.0
        _ = v4
        _ = v5
    }
    fmt.Println(v3)
    v6 := []int{}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int8 = -1
        v8 -= (v8 >> ((91) & 63))
        _ = fi7
        _ = v8
    }
    v3 = (v3 ^ v3)
    v9 := GenS1{F0: 289, F1: -1}
    var v10 uint = 2
    v2 += 816
    v11 := GenS2{F0: (v3 + 42), F1: v3}
    var v12 uint8 = 128
    var v13 int16 = int16(v2)
    if false {
        var v14 uint64 = 608
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 float32 = 0.0
            fmt.Println(v13)
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        var v17 uint16 = (^(v2))
        v12 -= v12
        _ = v17
    }
    fmt.Println(v3, v3)
    v10 += v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
