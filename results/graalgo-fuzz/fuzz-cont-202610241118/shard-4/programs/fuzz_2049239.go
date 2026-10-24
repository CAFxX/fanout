package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    var v4 uint8 = 128
    var v5 uint32 = 2147483647
    var v6 float64 = (0.0 / 3.14159)
    v4 = (v4 * (v4 % ((v4) | 1)))
    var v7 int32 = 1
    v6 -= v6
    v8 := []int{4294967295, 32768}
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int = (fi9 % ((fi9) | 1))
        _ = fi9
        _ = v10
    }
    var v11 uint8 = (^(v4))
    var v12 int = -122
    v13 := GenS1{F0: (^(v11)), F1: (false || false)}
    var v14 uint16 = 952
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS1{F0: 255, F1: (v5 == v5)}
        v5 -= v5
        for fi17 := 0; fi17 < 1; fi17++ {
            v4 += v11
            v7 = (v7 >> ((4294967295) & 63))
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    v4 = (v4 << ((uint(v1)) & 63))
    var v18 int = -9223372036854775808
    v5 = (^(v5))
    v5 = v5
    delete(v2, v5)
    var v19 int16 = (v3 | v3)
    var v20 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
