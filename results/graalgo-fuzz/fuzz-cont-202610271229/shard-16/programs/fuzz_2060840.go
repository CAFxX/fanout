package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    v1 := []uint32{1, 0, 280, 2147483647}
    var v2 uint16 = 255
    var v3 int8 = 0
    var v4 float64 = -2.25
    for fi5 := 0; fi5 < 8; fi5++ {
        v2 += v2
        _ = fi5
    }
    fmt.Println(v4)
    v6 := GenS1{F0: 255, F1: -9223372036854775808}
    fmt.Println(v2, v2, v4)
    var v7 uint64 = 0
    v2 = (v2 << ((794) & 63))
    var v8 int = -1
    var v9 float32 = (0.0 - 3.14159)
    v8 = v8
    var v10 uint64 = (v7 - (v7 | v7))
    var v11 int16 = -554
    v2 += (v2 << ((42) & 63))
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v1 = append(v1, 2147483647)
        _ = len(v1)
    } else {
        var v12 bool = ((!false) || true)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v13 := GenS1{F0: 251, F1: (v8 | v8)}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 104, F1: (^(v8))}
    var v15 uint64 = v10
    if (!false) {
        v10 += 453
    } else {
        var v16 float64 = ((1.5 / v4) / 0.1)
        _ = v16
    }
    var v17 uint64 = (v15 << ((2) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
