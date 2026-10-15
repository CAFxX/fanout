package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 bool }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = float64(3.14159)
    var v2 int16 = 1
    var v3 int16 = v2
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]uint32)
    v5[843] = 65535
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint8 = 43
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 127
    switch "hello" {
    case "hello":
        v1 += v1
    default:
        v9 := v4[uint32(v8)]
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS1{F0: true, F1: 65535, F2: (!(!false))}
        v12 := []int{9223372036854775807, 32767}
        fmt.Println(v2)
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: ("" + "")}
    var v14 int64 = -202
    v8 += (v8 ^ (v8 & v8))
    var v15 float32 = (0.0 + float32(1e10))
    v8 = (v8 - (v8 / (((^(v8))) | 1)))
    if (true || (42 <= 119)) {
        fmt.Println(v2)
        v16 := v5[479]
        var v17 int = 2147483647
        _ = v16
        _ = v17
    } else {
        var v18 int16 = (v2 + 32767)
        fmt.Println(v2, v3, v3)
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
