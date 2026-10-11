package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint }

func main() {
    var v1 int8 = -128
    var v2 int = 856
    v3 := []int{v2, (v2 % ((850) | 1)), (^(v2)), v2}
    var v4 float64 = 1.5
    var v5 int8 = (v1 & 123)
    var v6 int64 = -1
    if len(v3) > 0 {
        v3[0] = 128
    }
    var v7 float32 = (0.0 + 0.1)
    v4 = ((1.5 / (v4 - 1e10)) + (v4 / v4))
    var v8 int64 = v6
    switch 4294967295 {
    case 535:
        var v9 int16 = -195
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: 134, F2: 255}
    var v11 uint16 = 0
    if len(v3) > 0 {
        v3[0] = v2
    }
    v8 = 0
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint16 = (v11 << ((1) & 63))
        v14 := GenS1{F0: -32768, F1: 1, F2: 65535}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int32 = 958
    var v16 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
