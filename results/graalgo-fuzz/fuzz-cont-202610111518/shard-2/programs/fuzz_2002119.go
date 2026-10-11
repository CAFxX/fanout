package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }

func main() {
    var v1 int64 = 0
    v2 := []int{-692, 0, 1, 9223372036854775807}
    var v3 uint = 0
    v4 := make(map[uint32]uint32)
    v1 -= (v1 | -236)
    if (!true) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v5 int8 = 0
    var v6 uint8 = 127
    v7 := GenS1{F0: int8(v1), F1: 128}
    var v8 uint = v3
    v8 = v3
    if (v8 <= (v3 + v3)) {
        v5 += (v5 << ((796) & 63))
    } else {
        if (!true) {
            v9 := GenS1{F0: v5, F1: (^(v6))}
            var v10 uint8 = v6
            _ = v10
            _ = v9
        }
        var v11 int16 = 921
        v12 := GenS1{F0: (v5 / ((-1) | 1)), F1: (v6 ^ (v6 * (v6 | v6)))}
        _ = v11
        _ = v12
    }
    _ = len(v4)
    fmt.Println(v3, v6, v6)
    if ((^(v8)) <= (v3 + v3)) {
        v4[670] = 2
    }
    var v13 uint16 = 405
    var v14 uint64 = 1
    var v15 float64 = 0.1
    var v16 float32 = 1.5
    var v17 int8 = (v5 + (-(v5)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
