package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }

func main() {
    var v1 float64 = (-2.25 / 1.0)
    var v2 float64 = (0.1 / v1)
    v3 := make(map[uint32]int)
    var v4 bool = false
    var v5 uint16 = 32767
    v6 := []int{127, -2147483648, 127, 0}
    v4 = (2 != 2)
    var v7 uint64 = 42
    v8 := GenS1{F0: -1, F1: 128}
    v2 += 1e10
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint16 = v5
        v10 = (v5 << ((255) & 63))
        _ = fi9
        _ = v10
    }
    if len(v6) > 0 {
        v6[0] = 404
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int8 = 62
        if len(v6) > 0 {
            v6[0] = (fi11 ^ fi11)
        }
        _ = fi11
        _ = v12
    }
    var v13 bool = v4
    var v14 float64 = v1
    if v13 {
        v6 = append(v6, -9223372036854775808)
        v15 := GenS2{F0: 128, F1: (v5 - v5), F2: 510}
        _ = v15
    } else {
        var v16 float64 = (3.14159 + v14)
        _ = v16
    }
    delete(v3, 2)
    var v17 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
