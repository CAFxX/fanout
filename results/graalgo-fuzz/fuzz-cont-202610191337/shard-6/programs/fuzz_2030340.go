package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int32 = -592
    v4 := make(map[uint32]string)
    var v5 uint8 = 0
    switch 0 {
    case -258:
        v6 := GenS1{F0: -263, F1: -32768}
        _ = v6
    case 2147483647:
        v5 += (^(v5))
    case 32767:
        fmt.Println(v3, v3)
    default:
        var v7 uint32 = 2147483648
        _ = v7
    }
    var v8 int32 = (^(v3))
    if false {
        var v9 float64 = ((0.0 + float64(-2.25)) * 1e10)
        v10 := GenS1{F0: 650, F1: -793}
        v11 := GenS1{F0: 32768, F1: 42}
        _ = v10
        _ = v11
        _ = v9
    }
    v5 = 42
    var v12 int64 = -1
    var v13 bool = ((998 <= 0) && false)
    v3 = (v8 >> ((65535) & 63))
    var v14 bool = v13
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    fmt.Println(v8)
    if (!true) {
        var v16 int32 = (v3 & v8)
        fmt.Println(v8, v13, v14)
        var v17 int8 = 1
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 128, F1: 32767}
    v19 := GenS1{F0: 1, F1: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
