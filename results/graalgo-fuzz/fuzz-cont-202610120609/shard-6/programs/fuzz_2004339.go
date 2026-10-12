package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int64; F1 int8; F2 uint8 }
func genFn0(p0 bool, p1 uint8, p2 bool) int16 {
    var v6 float32 = (1.5 + 3.14159)
    var v7 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1() int16 {
    v8 := GenS1{F0: 42}
    var v9 int8 = -81
    _ = v8
    _ = v9
    return -32768
}

func main() {
    var v1 bool = false
    var v2 uint16 = 1
    var v3 uint64 = 4294967295
    var v4 uint32 = uint32(v2)
    v5 := []uint32{}
    var v10 uint8 = 136
    v5 = append(v5, (v4 + (v4 / ((v4) | 1))))
    v11 := GenS2{F0: 9223372036854775807, F1: 306}
    v12 := GenS3{F0: 0, F1: -128, F2: v10}
    for fi13 := 0; fi13 < 15; fi13++ {
        _ = len(v5)
        v14 := GenS2{F0: (fi13 >> ((2) & 63)), F1: 310}
        v1 = true
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        _ = len(v5)
        _ = fi15
    }
    v4 += v4
    v16 := GenS2{F0: 9223372036854775807, F1: int64(v10)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
