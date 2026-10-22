package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    var v1 uint8 = 128
    var v2 uint = 4294967295
    var v3 float64 = 1e10
    v4 := []uint32{}
    v4 = append(v4, 313)
    v5 := GenS1{F0: true, F1: v2}
    for fi6 := 0; fi6 < 13; fi6++ {
        v1 = v1
        v1 = 127
        _ = fi6
    }
    var v7 uint16 = 2
    var v8 int16 = 1000
    if (-2.25 == -2.25) {
        var v9 float32 = 1e10
        _ = v9
    } else {
        var v10 uint = 18446744073709551615
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int16 = (v8 % ((42) | 1))
        _ = fi11
        _ = v12
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v7 = v7
    v1 += v1
    _ = len(v4)
    var v13 int16 = -771
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
